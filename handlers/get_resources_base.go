package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getResourcesBaseHandler[T any](collection *mongo.Collection, sortOptions bson.M) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resources []T

		findOptions := options.Find().SetSort(sortOptions)
		cursor, err := collection.Find(c.Request.Context(), bson.M{}, findOptions)
		if err != nil {
			log.Println(err)
			respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
			return
		}
		defer cursor.Close(c.Request.Context())

		if err := cursor.All(c.Request.Context(), &resources); err != nil {
			log.Println("[ERROR] failed to decode records", err)
			respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
			return
		}

		respond(c, http.StatusOK, resources, nil)
	}
}
