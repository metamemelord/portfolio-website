package worker

import (
	"context"
	"errors"
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

var proxyRouteMutex sync.Mutex
var proxyRoutes map[string]model.ProxyItem
var proxyItemsCollection *mongo.Collection

func init() {
	proxyItemsCollection = db.GetCollection("proxy-items")
	proxyRoutes = make(map[string]model.ProxyItem)

	expiredProxyItemIDs := []primitive.ObjectID{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	cursor, err := proxyItemsCollection.Find(ctx, bson.M{"active": true})
	if err != nil {
		log.Println("Error while retrieving proxy items", err)
		return
	}
	proxyRouteMutex.Lock()
	for cursor.Next(ctx) {
		proxyItem := model.ProxyItem{}
		_ = cursor.Decode(&proxyItem)
		if time.Now().UnixNano() > proxyItem.Expiry.UnixNano() {
			expiredProxyItemIDs = append(expiredProxyItemIDs, proxyItem.ID)
		} else {
			proxyRoutes[proxyItem.RoutingKey] = proxyItem
		}
	}
	cursor.Close(ctx)
	proxyRouteMutex.Unlock()
	if len(expiredProxyItemIDs) > 0 {
		markInactive(&expiredProxyItemIDs)
	}
}

func ResolveProxyItem(routingKey string) (string, int, error) {
	routingKey = strings.ToLower(routingKey)
	var proxyItem model.ProxyItem
	var ok bool

	if proxyItem, ok = proxyRoutes[routingKey]; !ok {
		return "", 0, errors.New("Failed to find route")
	}

	if time.Now().UnixNano() > proxyItem.Expiry.UnixNano() {
		_ = DeleteProxyItem(routingKey)
		return "", 0, errors.New("Failed to find route")
	}

	statusCode := http.StatusTemporaryRedirect
	if proxyItem.Permanent {
		statusCode = http.StatusMovedPermanently
	}

	return proxyItem.Target, statusCode, nil
}

func AddProxyItem(ctx context.Context, proxyItem *model.ProxyItem) (string, error) {
	// Pre-processing
	proxyItem.ID = primitive.NewObjectID()
	proxyItem.RoutingKey = strings.ToLower(proxyItem.RoutingKey)
	exp, err := time.Parse(core.DATE_FORMAT, proxyItem.ExpiryString)
	if err != nil {
		proxyItem.Expiry = time.Now().Add(time.Hour * 876000)
	} else {
		proxyItem.Expiry = exp
	}

	result, err := proxyItemsCollection.InsertOne(ctx, proxyItem)
	if err != nil {
		log.Println(result, err)
		return "", err
	}
	proxyRouteMutex.Lock()
	proxyRoutes[proxyItem.RoutingKey] = *proxyItem
	proxyRouteMutex.Unlock()
	return proxyItem.ID.String(), err
}

func DeleteProxyItem(routingKey string) error {
	routingKey = strings.ToLower(routingKey)
	proxyItem, ok := proxyRoutes[routingKey]
	if !ok {
		return nil
	}

	proxyRouteMutex.Lock()
	delete(proxyRoutes, routingKey)
	proxyRouteMutex.Unlock()
	markInactive(&[]primitive.ObjectID{proxyItem.ID})
	return nil
}

func CheckAndMarkProxyInactive() {
	expiredProxyItemIDs := []primitive.ObjectID{}
	keysToBeDeleted := []string{}
	proxyRouteMutex.Lock()
	for k, proxyItem := range proxyRoutes {
		if time.Now().UnixNano() > proxyItem.Expiry.UnixNano() {
			expiredProxyItemIDs = append(expiredProxyItemIDs, proxyItem.ID)
			keysToBeDeleted = append(keysToBeDeleted, k)
		}
	}

	if len(keysToBeDeleted) > 0 {
		for _, k := range keysToBeDeleted {
			delete(proxyRoutes, k)
		}
	}
	proxyRouteMutex.Unlock()

	if len(expiredProxyItemIDs) > 0 {
		markInactive(&expiredProxyItemIDs)
	}
}

func markInactive(expiredProxyItemIDs *[]primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	log.Println("Marking inactive IDs")
	log.Println(proxyItemsCollection.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": expiredProxyItemIDs}}, bson.M{"$set": bson.M{"active": false}}))
}
