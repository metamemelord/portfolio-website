package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/metamemelord/portfolio-website/db"
	"github.com/metamemelord/portfolio-website/model"
	"github.com/metamemelord/portfolio-website/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var redirectionRouteMutex sync.Mutex
var redirectionRoutes map[string]model.RedirectionItem
var redirectionItemsCollection *mongo.Collection
var redirectionHitCounterChannel = make(chan primitive.ObjectID, 200)

func init() {
	redirectionItemsCollection = db.GetCollection("redirection-items")
	redirectionRoutes = make(map[string]model.RedirectionItem)

	expiredRedirectionItemIDs := []primitive.ObjectID{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	cursor, err := redirectionItemsCollection.Find(ctx, bson.M{"active": true})
	if err != nil {
		log.Println("Error while retrieving redirection items", err)
		return
	}
	redirectionRouteMutex.Lock()
	for cursor.Next(ctx) {
		redirectionItem := model.RedirectionItem{}
		_ = cursor.Decode(&redirectionItem)
		if time.Now().UTC().UnixNano() > redirectionItem.Expiry.UnixNano() {
			expiredRedirectionItemIDs = append(expiredRedirectionItemIDs, redirectionItem.ID)
		} else {
			redirectionRoutes[redirectionItem.RoutingKey] = redirectionItem
		}
	}
	cursor.Close(ctx)
	redirectionRouteMutex.Unlock()
	if len(expiredRedirectionItemIDs) > 0 {
		markInactive(&expiredRedirectionItemIDs)
	}

	go runRedirectionHitCounterPublisher(context.Background())
}

func ResolveRedirectionItem(routingKey, pathToForward, rawQuery string) (string, int, error) {
	routingKey = strings.ToLower(routingKey)
	var redirectionItem model.RedirectionItem
	var ok bool
	var rawQueryUsed bool

	if redirectionItem, ok = redirectionRoutes[routingKey]; !ok {
		return core.EMPTY_STRING, 0, errors.New("Failed to find route")
	}

	if time.Now().UTC().UnixNano() > redirectionItem.Expiry.UnixNano() {
		_ = DeleteRedirectionItem(routingKey)
		return core.EMPTY_STRING, 0, errors.New("Route has been expired")
	}

	statusCode := http.StatusTemporaryRedirect
	if redirectionItem.Permanent {
		statusCode = http.StatusMovedPermanently
	}

	target := redirectionItem.Target

	if *redirectionItem.ForwardPath {
		target = fmt.Sprintf("%s/%s", strings.TrimRight(target, "/"), strings.TrimLeft(pathToForward, "/"))
	}

	if len(rawQuery) > 0 {
		target = fmt.Sprintf("%s?%s", target, rawQuery)
		rawQueryUsed = true
	}

	// Process metadata here
	if redirectionItem.Metadata != nil {
		for metadataItemType, data := range redirectionItem.Metadata {
			switch metadataItemType {
			case model.MetadataItemTypeQueryParam.String():
				queries := make([]string, len(data))
				iter := 0
				for k, v := range data {
					queries[iter] = fmt.Sprintf("%s=%s", k, v.(string))
					iter++
				}
				joiner := core.QUESTION_MARK
				if rawQueryUsed {
					joiner = core.AMPERSAND
				}
				target = fmt.Sprintf("%s%s%s", target, joiner, strings.Join(queries, core.AMPERSAND))
			default:
				break
			}
		}
	}

	log.Printf("Redirecting to %s\n", target)

	go func(id primitive.ObjectID) {
		redirectionHitCounterChannel <- id
	}(redirectionItem.ID)

	return target, statusCode, nil
}

func AddRedirectionItem(ctx context.Context, redirectionItem *model.RedirectionItem) (string, error) {
	// Pre-processing
	redirectionItem.ID = primitive.NewObjectID()
	redirectionItem.RoutingKey = strings.ToLower(redirectionItem.RoutingKey)

	target, err := url.Parse(redirectionItem.Target)
	if err != nil {
		log.Println("Invalid path:", redirectionItem.Target, err)
		return core.EMPTY_STRING, err
	}

	if target.Scheme == core.EMPTY_STRING {
		redirectionItem.Target = fmt.Sprintf("https://%s", redirectionItem.Target)
	} else if target.Scheme != "https" && target.Scheme != "http" {
		msg := "Currently only http and https redirects are supported"
		log.Println(msg, ":", redirectionItem.Target)
		return core.EMPTY_STRING, errors.New(msg)
	}

	exp, err := time.Parse(core.DATE_FORMAT, redirectionItem.ExpiryString)
	if err != nil {
		redirectionItem.Expiry = time.Now().UTC().Add(time.Hour * 876000)
		redirectionItem.Permanent = true
	} else {
		redirectionItem.Expiry = exp
		redirectionItem.Permanent = false
	}

	// Cleaning up unintended addition of defaults
	redirectionItem.ExpiryString = core.EMPTY_STRING
	redirectionItem.HitCount = 0

	redirectionItem.Active = time.Now().UTC().UnixNano() < redirectionItem.Expiry.UnixNano()
	if redirectionItem.ForwardPath == nil {
		forwardPath := true
		redirectionItem.ForwardPath = &forwardPath
	}

	result, err := redirectionItemsCollection.InsertOne(ctx, redirectionItem)
	if err != nil {
		log.Println(result, err)
		return core.EMPTY_STRING, err
	}
	redirectionRouteMutex.Lock()
	redirectionRoutes[redirectionItem.RoutingKey] = *redirectionItem
	redirectionRouteMutex.Unlock()
	return redirectionItem.RoutingKey, err
}

func DeleteRedirectionItem(routingKey string) error {
	routingKey = strings.ToLower(routingKey)
	redirectionItem, ok := redirectionRoutes[routingKey]
	if !ok {
		return nil
	}

	redirectionRouteMutex.Lock()
	delete(redirectionRoutes, routingKey)
	redirectionRouteMutex.Unlock()
	markInactive(&[]primitive.ObjectID{redirectionItem.ID})
	return nil
}

func CheckAndMarkRedirectionInactive() {
	expiredRedirectionItemIDs := []primitive.ObjectID{}
	keysToBeDeleted := []string{}
	redirectionRouteMutex.Lock()
	for k, redirectionItem := range redirectionRoutes {
		if time.Now().UTC().UnixNano() > redirectionItem.Expiry.UnixNano() {
			expiredRedirectionItemIDs = append(expiredRedirectionItemIDs, redirectionItem.ID)
			keysToBeDeleted = append(keysToBeDeleted, k)
		}
	}

	if len(keysToBeDeleted) > 0 {
		for _, k := range keysToBeDeleted {
			delete(redirectionRoutes, k)
		}
	}
	redirectionRouteMutex.Unlock()

	if len(expiredRedirectionItemIDs) > 0 {
		markInactive(&expiredRedirectionItemIDs)
	}
}

func markInactive(expiredRedirectionItemIDs *[]primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	log.Println("Marking inactive IDs")
	log.Println(redirectionItemsCollection.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": expiredRedirectionItemIDs}}, bson.M{"$set": bson.M{"active": false}}))
}

func runRedirectionHitCounterPublisher(ctx context.Context) {
	for {
		select {
		case value := <-redirectionHitCounterChannel:
			go incrementHitCount(ctx, value)
		case <-ctx.Done():
			break
		}
	}
}

func incrementHitCount(ctx context.Context, id primitive.ObjectID) {
	result, err := redirectionItemsCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$inc": bson.M{"hit_count": 1}})
	if err == nil && (result.MatchedCount != 1 || result.ModifiedCount != 1) {
		err = fmt.Errorf("Error incrementing the hit count Matched=%d, Modified=%d", result.MatchedCount, result.ModifiedCount)
	}
	log.Printf("Updating hit count for id=(%s), error=(%v)", id.String(), err)
}
