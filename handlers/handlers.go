package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/db"
	"github.com/metamemelord/portfolio-website/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	blogPostCollection    *mongo.Collection
	ErrParseRequestBody   = fmt.Errorf("Failed to parse request body")
	ErrInternalServer     = fmt.Errorf("Internal server error")
	ErrPostNotFoundWithID = fmt.Errorf("Could not find a post with that _id")
)

func init() {
	blogPostCollection = db.BlogPostCollection()
}

func Register(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.GET("/blogs", returnBlogPosts)
		api.POST("/blog", addBlogPost)
		api.PUT("/blog", updateBlogPost)
		api.DELETE("/blog", deleteBlogPost)
	}
	g.Static("/js", "./dist/js")
	g.Static("/css", "./dist/css")
	g.Static("/img", "./dist/img")
	g.StaticFile("/favicon.ico", "./dist/favicon.ico")
	g.NoRoute(htmlSupplier)
}

func returnBlogPosts(c *gin.Context) {
	blogPosts := []*model.BlogPost{}
	cursor, err := blogPostCollection.Find(c.Request.Context(), bson.M{"visible": true})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, ErrInternalServer)
		return
	}

	for cursor.Next(c.Request.Context()) {
		blogPost := &model.BlogPost{}
		err = cursor.Decode(blogPost)
		blogPosts = append(blogPosts, blogPost)
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, blogPosts)
}

func addBlogPost(c *gin.Context) {
	blogPost := &model.BlogPost{}

	if err := c.ShouldBindJSON(blogPost); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, ErrParseRequestBody)
		return
	}

	blogPost.ID = primitive.NewObjectID()
	res, err := blogPostCollection.InsertOne(c.Request.Context(), blogPost)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func updateBlogPost(c *gin.Context) {
	body := make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&body)
	ID, ok := body["_id"]
	if !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("_id cannot be empty"))
	}

	delete(body, "_id")

	res, err := blogPostCollection.UpdateOne(c.Request.Context(), bson.M{"_id": ID}, body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if res.ModifiedCount == 0 {
		c.AbortWithError(http.StatusNotFound, ErrPostNotFoundWithID)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func deleteBlogPost(c *gin.Context) {
	body := make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&body)
	ID, ok := body["_id"]
	if !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("_id cannot be empty"))
	}

	res, err := blogPostCollection.UpdateOne(c.Request.Context(), bson.M{"visible": true, "_id": ID}, bson.M{"visible": false})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if res.ModifiedCount == 0 {
		c.AbortWithError(http.StatusNotFound, ErrPostNotFoundWithID)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func htmlSupplier(c *gin.Context) {
	file, _ := ioutil.ReadFile("./dist/index.html")
	c.Data(200, "text/html", file)
}
