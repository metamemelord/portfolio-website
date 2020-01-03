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

func getExperiences(c *gin.Context) {
	experiences := []*model.Experience{}
	findOptions := options.Find().SetSort(bson.M{"_id": -1})
	cursor, err := experiencesCollection.Find(c.Request.Context(), bson.M{}, findOptions)
	if err != nil {
		log.Println(err)
		respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
		return
	}

	for cursor.Next(c.Request.Context()) {
		experience := &model.Experience{}
		err = cursor.Decode(experience)
		experiences = append(experiences, experience)
	}

	if err != nil {
		respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
		return
	}

	respond(c, http.StatusOK, experiences, nil)
}

func addExperience(c *gin.Context) {
	experience := &model.Experience{}

	if err := c.ShouldBindJSON(experience); err != nil {
		respond(c, http.StatusBadRequest, nil, ErrParseRequestBody)
		return
	}

	exp := experiencesCollection.FindOne(c.Request.Context(), bson.M{"company": experience.Company, "title": experience.Title})
	if exp.Err() == nil {
		respond(c, http.StatusConflict, nil, fmt.Errorf("Entry already exists"))
		return
	}

	experience.ID = primitive.NewObjectID()
	res, err := experiencesCollection.InsertOne(c.Request.Context(), experience)
	if err != nil {
		respond(c, http.StatusInternalServerError, nil, err)
		return
	}

	respond(c, http.StatusCreated, res, nil)
}

func updateExperience(c *gin.Context) {
	experience := &model.Experience{}

	if err := c.ShouldBindJSON(experience); err != nil {
		respond(c, http.StatusBadRequest, nil, ErrParseRequestBody)
		return
	}

	exp := experiencesCollection.FindOneAndUpdate(c.Request.Context(), bson.M{"company": experience.Company, "title": experience.Title}, experience)
	if exp.Err() != nil {
		respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Failed to update exprience"))
		return
	}

	respond(c, http.StatusCreated, exp, nil)
}

func deleteExperience(c *gin.Context) {
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
		exp := experiencesCollection.FindOneAndDelete(c.Request.Context(), bson.M{"_id": dID})
		if exp.Err() != nil {
			log.Println(exp.Err())
			respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Failed to delete exprience"))
			return
		}
	}
	respond(c, http.StatusOK, nil, nil)
}
