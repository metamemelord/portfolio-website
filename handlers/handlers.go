package handlers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	blogPostCollection    *mongo.Collection
	experiencesCollection *mongo.Collection
	ErrParseRequestBody   = fmt.Errorf("Failed to parse request body")
	ErrInternalServer     = fmt.Errorf("Internal server error")
	ErrPostNotFoundWithID = fmt.Errorf("Could not find a post with that _id")
)

func init() {
	blogPostCollection = db.GetCollection("blog-posts")
	experiencesCollection = db.GetCollection("experiences")
}

func Register(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.GET("/blogs", returnBlogPosts)
		api.POST("/blog", verifyCredentials, addBlogPost)
		api.PUT("/blog", verifyCredentials, updateBlogPost)
		api.DELETE("/blog", verifyCredentials, deleteBlogPost)

		api.GET("/experiences", getExperiences)
		api.POST("/experience", verifyCredentials, addExperience)
		api.PUT("/experience", verifyCredentials, updateExperience)
		api.DELETE("/experience", verifyCredentials, deleteExperience)

		api.GET("/repos", getGithubReposHandler)
		api.GET("/wordpress", getWordpressPostsHandler)
		api.POST("admin/data/refresh", verifyCredentials, refreshData)
	}
	g.GET("/health", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})
	g.Static("/js", "./dist/js")
	g.Static("/css", "./dist/css")
	g.Static("/img", "./dist/img")
	g.StaticFile("/favicon.ico", "./dist/favicon.ico")
	g.NoRoute(htmlSupplier)
}

func verifyCredentials(c *gin.Context) {
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
