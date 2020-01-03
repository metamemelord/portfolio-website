package db

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
var once sync.Once

func getDatabase() *mongo.Database {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
		mongoClient, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalln("Could not connect to mongo:", err)
		}
		database = mongoClient.Database("portfolio")
	})
	return database
}

func GetCollection(collectionName string) *mongo.Collection {
	return getDatabase().Collection(collectionName)
}
