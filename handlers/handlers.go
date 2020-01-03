package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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
		api.POST("/blog", verifyCredentuals, addBlogPost)
		api.PUT("/blog", verifyCredentuals, updateBlogPost)
		api.DELETE("/blog", verifyCredentuals, deleteBlogPost)
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

func verifyCredentuals(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		auth = c.GetHeader("authorization")
	}

	if auth == "" {
		respond(c, http.StatusUnauthorized, nil, fmt.Errorf("Require auth"))
		c.Abort()
		return
	}

	authStringTokens := strings.Split(auth, " ")
	if len(authStringTokens) != 2 && strings.ToLower(authStringTokens[0]) != "basic" {
		respond(c, http.StatusUnauthorized, nil, fmt.Errorf("Invalid auth"))
		c.Abort()
		return
	}

	encodedUsernamePassword := authStringTokens[1]
	usernamePasswordString, err := base64.StdEncoding.DecodeString(encodedUsernamePassword)

	if err != nil || os.Getenv("APP_AUTH") != string(usernamePasswordString) {
		respond(c, http.StatusUnauthorized, nil, fmt.Errorf("Invalid username or password"))
		c.Abort()
		return
	}
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

	res, err := blogPostCollection.UpdateOne(c.Request.Context(), bson.M{"_id": ID}, body)

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

	res, err := blogPostCollection.UpdateOne(c.Request.Context(), bson.M{"visible": true, "_id": ID}, bson.M{"visible": false})

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

func htmlSupplier(c *gin.Context) {
	file, _ := ioutil.ReadFile("./dist/index.html")
	c.Data(200, "text/html", file)
}

func respond(c *gin.Context, status int, payload interface{}, err error) {
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(status, map[string]interface{}{"error": err.Error()})
	} else {
		log.Println("[INFO]: ", payload)
		c.JSON(status, payload)
	}
}
