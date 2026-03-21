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

// returnBlogPosts retrieves all visible blog posts
// @Summary Get all blog posts
// @Description Get all visible blog posts
// @Accept json
// @Produce json
// @Success 200 {array} model.BlogPost
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blogs [get]
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

// addBlogPost creates a new blog post
// @Summary Add a new blog post
// @Description Create a new blog post (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param blogPost body model.BlogPost true "Blog post data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 409 {object} map[string]interface{} "Conflict"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog [post]
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

// updateBlogPost updates an existing blog post
// @Summary Update a blog post
// @Description Update an existing blog post (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param update body map[string]interface{} true "Blog post updates (must include _id)"
// @Success 200
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Post not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog [put]
func updateBlogPost(c *gin.Context) {
	body := make(map[string]string)
	_ = json.NewDecoder(c.Request.Body).Decode(&body)
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

// deleteBlogPost soft-deletes a blog post
// @Summary Delete a blog post
// @Description Soft-delete a blog post by marking it as invisible (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param delete body map[string]interface{} true "Post ID (_id required)"
// @Success 200
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Post not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog [delete]
func deleteBlogPost(c *gin.Context) {
	body := make(map[string]string)
	_ = json.NewDecoder(c.Request.Body).Decode(&body)
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
