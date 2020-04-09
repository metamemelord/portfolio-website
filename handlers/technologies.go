package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getTechnologies(c *gin.Context) {
	technologies := []*model.Technology{}
	findOptions := options.Find().SetSort(bson.M{"order": 1})
	cursor, err := technologiesCollection.Find(c.Request.Context(), bson.M{}, findOptions)
	if err != nil {
		log.Println(err)
		respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
		return
	}

	for cursor.Next(c.Request.Context()) {
		technology := &model.Technology{}
		err = cursor.Decode(technology)
		technologies = append(technologies, technology)
	}

	if err != nil {
		respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
		return
	}

	respond(c, http.StatusOK, technologies, nil)
}

func addTechnology(c *gin.Context) {
	technology := &model.Technology{}

	if err := c.ShouldBindJSON(technology); err != nil {
		respond(c, http.StatusBadRequest, nil, ErrParseRequestBody)
		return
	}

	exp := technologiesCollection.FindOne(c.Request.Context(), bson.M{"name": technology.Name})
	if exp.Err() == nil {
		respond(c, http.StatusConflict, nil, fmt.Errorf("Entry already exists"))
		return
	}

	technology.ID = primitive.NewObjectID()
	res, err := technologiesCollection.InsertOne(c.Request.Context(), technology)
	if err != nil {
		respond(c, http.StatusInternalServerError, nil, err)
		return
	}

	respond(c, http.StatusCreated, res, nil)
}

func updateTechnology(c *gin.Context) {
	technology := &model.Technology{}

	if err := c.ShouldBindJSON(technology); err != nil {
		respond(c, http.StatusBadRequest, nil, ErrParseRequestBody)
		return
	}

	exp := technologiesCollection.FindOneAndUpdate(c.Request.Context(), bson.M{"name": technology.Name}, bson.M{"$set": technology})
	if exp.Err() != nil {
		log.Println(exp.Err())
		respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Failed to update technology"))
		return
	}

	respond(c, http.StatusOK, technology, nil)
}

func deleteTechnology(c *gin.Context) {
	expMap := make(map[string]string)

	if err := json.NewDecoder(c.Request.Body).Decode(&expMap); err != nil {
		respond(c, http.StatusBadRequest, nil, ErrParseRequestBody)
		return
	}
	if id, ok := expMap["_id"]; !ok {
		respond(c, http.StatusBadRequest, nil, fmt.Errorf("_id cannot be empty"))
		return
	} else {
		dID, _ := primitive.ObjectIDFromHex(id)
		exp := technologiesCollection.FindOneAndDelete(c.Request.Context(), bson.M{"_id": dID})
		if exp.Err() != nil {
			log.Println(exp.Err())
			respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Failed to delete technology"))
			return
		}
	}
	respond(c, http.StatusOK, nil, nil)
}
