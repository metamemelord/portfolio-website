package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/db"
	"github.com/metamemelord/portfolio-website/pkg/core"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	blogPostCollection     *mongo.Collection
	experiencesCollection  *mongo.Collection
	technologiesCollection *mongo.Collection
	staticAssetsBasePath   = "/srv/portfolio/dist"
	ErrParseRequestBody    = fmt.Errorf("Failed to parse request body")
	ErrInternalServer      = fmt.Errorf("Internal server error")
	ErrPostNotFoundWithID  = fmt.Errorf("Could not find a post with that _id")

	htmlBody = []byte{}
	err      error

	indexHtmlBasePath = staticAssetsBasePath + "/index.html"
	jsBasePath        = staticAssetsBasePath + "/js"
	cssBasePath       = staticAssetsBasePath + "/css"
	imgBasePath       = staticAssetsBasePath + "/img"
	faviconBasePath   = staticAssetsBasePath + "/favicon.ico"
	robotsTxtBasePath = staticAssetsBasePath + "/robots.txt"
	sitemapBasePath   = staticAssetsBasePath + "/sitemap.xml"
)

func init() {
	blogPostCollection = db.GetCollection("blog-posts")
	experiencesCollection = db.GetCollection("experiences")
	technologiesCollection = db.GetCollection("technologies")

	htmlBody, err = os.ReadFile(indexHtmlBasePath)
	if err != nil {
		log.Println("[ERROR]: Failed to load HTML (" + err.Error() + ")")
	}
}

func Register(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.GET("/blogs", cacheSetter(2*time.Hour), returnBlogPosts)
		api.POST("/blog", verifyCredentials, addBlogPost)
		api.PUT("/blog", verifyCredentials, updateBlogPost)
		api.DELETE("/blog", verifyCredentials, deleteBlogPost)

		api.GET("/experiences", cacheSetter(720*time.Hour), getExperiences)
		api.POST("/experience", verifyCredentials, addExperience)
		api.PUT("/experience", verifyCredentials, updateExperience)
		api.DELETE("/experience", verifyCredentials, deleteExperience)

		api.GET("/technologies", cacheSetter(720*time.Hour), getTechnologies)
		api.POST("/technology", verifyCredentials, addTechnology)
		api.PUT("/technology", verifyCredentials, updateTechnology)
		api.DELETE("/technology", verifyCredentials, deleteTechnology)

		api.GET("/repos", cacheSetter(2*time.Hour), getGithubReposHandler)
		api.GET("/wordpress", getWordpressPostsHandler)
		api.GET("/wordpress/:id", getWordpressPostbyIDHandler)
		api.POST("/email", sendEmail)
		api.POST("/admin/data/refresh", verifyCredentials, refreshData)
	}

	g.GET("/health", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNoContent)
	})

	// Redirection routes
	{
		g.GET("/r/:routing_key", resolveRedirection, htmlSupplier)
		g.GET("/r/:routing_key/*path", resolveRedirection, htmlSupplier)

		// API
		api.GET("/redirect", verifyCredentials, getRedirectionItems)
		api.GET("/redirect/:routing_key", verifyCredentials, getRedirectionItemByRoutingKey)
		api.POST("/redirect", verifyCredentials, addRedirection)
		api.DELETE("/redirect/:routing_key", verifyCredentials, deleteRedirection)
	}

	public := g.Group("/", cacheSetter(168*time.Hour))
	{
		public.Static("/js", jsBasePath)
		public.Static("/css", cssBasePath)
		public.Static("/img", imgBasePath)
		public.StaticFile("/favicon.ico", faviconBasePath)
		public.StaticFile("/robots.txt", robotsTxtBasePath)
		public.StaticFile("/sitemap.xml", sitemapBasePath)
	}

	g.NoRoute(htmlSupplier)
}

func verifyCredentials(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == core.EMPTY_STRING {
		auth = c.GetHeader("authorization")
	}

	if auth == core.EMPTY_STRING {
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
	c.Data(200, "text/html", htmlBody)
}

func respond(c *gin.Context, status int, payload interface{}, err error) {
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(status, map[string]interface{}{"error": err.Error()})
	} else {
		if payload != nil {
			resp, _ := json.Marshal(payload)
			log.Println("[INFO]: ", string(resp))
			c.Data(status, "application/json", resp)
		} else {
			log.Println("[INFO]: Status OK")
			c.String(status, core.EMPTY_STRING)
		}
	}
}

func cacheSetter(t time.Duration) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", fmt.Sprintf("max-age=%d, public", int64(t.Seconds())))
		c.Next()
	}
}
