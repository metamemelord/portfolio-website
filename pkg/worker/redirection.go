package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
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
			case model.MetadataItemTypeQueryParameter:
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

	return target, statusCode, nil
}

func AddRedirectionItem(ctx context.Context, redirectionItem *model.RedirectionItem) (string, error) {
	// Pre-processing
	redirectionItem.ID = primitive.NewObjectID()
	redirectionItem.RoutingKey = strings.ToLower(redirectionItem.RoutingKey)
	exp, err := time.Parse(core.DATE_FORMAT, redirectionItem.ExpiryString)
	if err != nil {
		redirectionItem.Expiry = time.Now().UTC().Add(time.Hour * 876000)
	} else {
		redirectionItem.Expiry = exp
	}
	redirectionItem.ExpiryString = core.EMPTY_STRING
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
