package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func returnBlogPosts(c *gin.Context) {
	blogPosts := []*model.BlogPost{}
	cursor, err := blogPostCollection.Find(c.Request.Context(), bson.M{"visible": true})
	if err != nil {
		respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
		return
	}

	for cursor.Next(c.Request.Context()) {
		blogPost := &model.BlogPost{}
		err = cursor.Decode(blogPost)
		blogPosts = append(blogPosts, blogPost)
	}

	if err != nil {
		respond(c, http.StatusInternalServerError, nil, ErrInternalServer)
		return
	}

	respond(c, http.StatusOK, blogPosts, nil)
}

func addBlogPost(c *gin.Context) {
	blogPost := &model.BlogPost{Tags: []string{}}

	if err := c.ShouldBindJSON(blogPost); err != nil {
		respond(c, http.StatusBadRequest, nil, ErrParseRequestBody)
		return
	}

	bp := blogPostCollection.FindOne(c.Request.Context(), bson.M{"title": blogPost.Title})
	if bp.Err() == nil {
		respond(c, http.StatusConflict, nil, fmt.Errorf("Entry with this title already exists"))
		return
	}

	blogPost.ID = primitive.NewObjectID()
	res, err := blogPostCollection.InsertOne(c.Request.Context(), blogPost)
	if err != nil {
		respond(c, http.StatusInternalServerError, nil, err)
		return
	}

	respond(c, http.StatusCreated, res, nil)
}

func updateBlogPost(c *gin.Context) {
	body := make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&body)
	ID, ok := body["_id"]
	if !ok {
		respond(c, http.StatusBadRequest, nil, fmt.Errorf("_id cannot be empty"))
		return
	}

	delete(body, "_id")

	bpID, _ := primitive.ObjectIDFromHex(ID)
	res, err := blogPostCollection.UpdateOne(c.Request.Context(), bson.M{"_id": bpID}, body)

	if err != nil {
		respond(c, http.StatusInternalServerError, nil, err)
		return
	}

	if res.ModifiedCount == 0 {
		respond(c, http.StatusNotFound, nil, ErrPostNotFoundWithID)
		return
	}

	respond(c, http.StatusOK, nil, nil)
}

func deleteBlogPost(c *gin.Context) {
	body := make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&body)
	ID, ok := body["_id"]
	if !ok {
		respond(c, http.StatusBadRequest, nil, fmt.Errorf("_id cannot be empty"))
		return
	}

	bpID, _ := primitive.ObjectIDFromHex(ID)
	res, err := blogPostCollection.UpdateOne(c.Request.Context(), bson.M{"visible": true, "_id": bpID}, bson.M{"visible": false})

	if err != nil {
		respond(c, http.StatusInternalServerError, nil, err)
		return
	}

	if res.ModifiedCount == 0 {
		respond(c, http.StatusNotFound, nil, ErrPostNotFoundWithID)
		return
	}

	respond(c, http.StatusOK, nil, nil)
}
