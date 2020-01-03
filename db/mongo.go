package db

import "go.mongodb.org/mongo-driver/mongo"

import "go.mongodb.org/mongo-driver/mongo/options"

import "os"

import "context"

import "log"

func BlogPostCollection() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalln("Could not connect to mongo:", err)
	}
	database := mongoClient.Database("portfolio")
	if database == nil {
		log.Fatalln("Database \"portfolio\" does not exist")
	}

	collection := database.Collection("blog-posts")
	if collection == nil {
		log.Fatalln("Database \"portfolio\" does not have \"blog-posts\" collection")
	}
	return collection
}
