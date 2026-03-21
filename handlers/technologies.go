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
)

// getTechnologies retrieves all technologies
// @Summary Get technologies
// @Description Get all technologies sorted by order
// @Accept json
// @Produce json
// @Success 200 {array} model.Technology
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /technologies [get]
func getTechnologies(c *gin.Context) {
	getResourcesBaseHandler[model.Technology](technologiesCollection, bson.M{"_id": 1})(c)
}

// addTechnology adds a new technology
// @Summary Add a new technology
// @Description Create a new technology entry (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param technology body model.Technology true "Technology data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 409 {object} map[string]interface{} "Entry already exists"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /technology [post]
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

// updateTechnology updates an existing technology
// @Summary Update a technology
// @Description Update an existing technology entry (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param technology body model.Technology true "Updated technology data"
// @Success 200 {object} model.Technology
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /technology [put]
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

// deleteTechnology deletes a technology
// @Summary Delete a technology
// @Description Delete a technology entry (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param delete body map[string]interface{} true "Technology ID (_id required)"
// @Success 200
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /technology [delete]
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
