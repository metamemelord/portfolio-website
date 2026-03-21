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
		resources := []*T{}
		findOptions := options.Find().SetSort(sortOptions)
		cursor, err := collection.Find(c.Request.Context(), bson.M{}, findOptions)
		if err != nil {
			log.Println(err)
			respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
			return
		}

		for cursor.Next(c.Request.Context()) {
			resource := new(T)
			err = cursor.Decode(resource)
			resources = append(resources, resource)
		}

		if err != nil {
			respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
			return
		}

		respond(c, http.StatusOK, resources, nil)
	}
}
