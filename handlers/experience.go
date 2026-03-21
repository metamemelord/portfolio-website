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

// getExperiences retrieves all work experiences
// @Summary Get experiences
// @Description Get all user work experiences sorted by most recent
// @Accept json
// @Produce json
// @Success 200 {array} model.Experience
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /experiences [get]
func getExperiences(c *gin.Context) {
	getResourcesBaseHandler[model.Experience](experiencesCollection, bson.M{"_id": -1})(c)
}

// addExperience adds a new work experience
// @Summary Add a new experience
// @Description Create a new work experience entry (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param experience body model.Experience true "Experience data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 409 {object} map[string]interface{} "Entry already exists"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /experience [post]
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

// updateExperience updates an existing work experience
// @Summary Update an experience
// @Description Update an existing work experience entry (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param experience body model.Experience true "Updated experience data"
// @Success 200 {object} model.Experience
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /experience [put]
func updateExperience(c *gin.Context) {
	experience := &model.Experience{}

	if err := c.ShouldBindJSON(experience); err != nil {
		respond(c, http.StatusBadRequest, nil, ErrParseRequestBody)
		return
	}

	exp := experiencesCollection.FindOneAndUpdate(c.Request.Context(), bson.M{"company": experience.Company, "title": experience.Title}, bson.M{"$set": experience})
	if exp.Err() != nil {
		respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Failed to update exprience"))
		return
	}

	respond(c, http.StatusOK, experience, nil)
}

// deleteExperience deletes a work experience
// @Summary Delete an experience
// @Description Delete a work experience entry (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param delete body map[string]interface{} true "Experience ID (_id required)"
// @Success 200
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /experience [delete]
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
