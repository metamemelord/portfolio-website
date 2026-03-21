package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/metamemelord/portfolio-website/model"
	"github.com/metamemelord/portfolio-website/pkg/communication"
	"github.com/metamemelord/portfolio-website/pkg/worker"
)

type UserProfile struct {
	ID             string         `json:"id" bson:"id"`
	Name           string         `json:"name" bson:"name"`
	Email          string         `json:"email" bson:"email"`
	Location       string         `json:"location" bson:"location"`
	PhoneNumber    string         `json:"phone_number" bson:"phone_number"`
	DynamicContent map[string]any `json:"dynamic_content" bson:"dynamic_content"`
	Occupation     Occupation     `json:"occupation" bson:"occupation"`
}

type Occupation struct {
	Title   string `json:"title" bson:"title"`
	Company string `json:"company" bson:"company"`
	Since   int    `json:"since" bson:"since"`
}

var emailClient communication.EmailClient
var userProfile *UserProfile
var loadProfileOnce sync.Once

func init() {
	emailClient = communication.NewMicrosoft365EmailClient()
	loadProfileOnce.Do(func() {
		go func() {
			time.Sleep(time.Second * 5)
			res := profileCollection.FindOne(context.Background(), bson.M{})
			if res.Err() != nil {
				log.Println("Failed to read profile data from mongo")
				return
			}
			err := res.Decode(userProfile)
			if err != nil {
				log.Println("Could not load profile data")
			}
		}()
	})
}

// getProfile retrieves user profile information
// @Summary Get user profile
// @Description Get the user's profile information
// @Accept json
// @Produce json
// @Success 200 {object} UserProfile
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /profile [get]
func getProfile(c *gin.Context) {
	respond(c, 200, userProfile, nil)
}

// sendEmail sends an email through the contact form
// @Summary Send email
// @Description Send an email message through the portfolio contact form
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Sender's name"
// @Param email formData string true "Sender's email"
// @Param datetime formData string true "Date and time of submission"
// @Param body formData string true "Email body/message"
// @Success 200
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /email [post]
func sendEmail(c *gin.Context) {
	email := &communication.Email{
		SenderName:  c.PostForm("name"),
		SenderEmail: c.PostForm("email"),
		DateTime:    c.PostForm("datetime"),
		Body:        c.PostForm("body"),
	}

	_, err := emailClient.Send(c.Request.Context(), email)
	if err != nil {
		log.Println("[ERROR] ", err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

// refreshData forces a refresh of portfolio data
// @Summary Refresh portfolio data
// @Description Force a refresh of all portfolio data from external sources
// @Accept json
// @Produce json
// @Security BasicAuth
// @Success 202
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /admin/data/refresh [post]
func refreshData(c *gin.Context) {
	log.Println("Force refresh called")
	worker.RefreshData()
	respond(c, http.StatusAccepted, nil, nil)
}

// getGithubReposHandler retrieves GitHub repositories
// @Summary Get GitHub repositories
// @Description Get all GitHub repositories from the user's GitHub profile
// @Accept json
// @Produce json
// @Success 200 {array} object
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /repos [get]
func getGithubReposHandler(c *gin.Context) {
	log.Println("[INFO] Github data accessed")
	c.JSON(http.StatusOK, worker.GetData().GithubData)
}

// getWordpressPostsHandler retrieves WordPress posts
// @Summary Get WordPress posts
// @Description Get all WordPress posts from the user's WordPress blog
// @Accept json
// @Produce json
// @Success 200 {array} object
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /wordpress [get]
func getWordpressPostsHandler(c *gin.Context) {
	log.Println("[INFO] Wordpress data accessed")
	c.JSON(http.StatusOK, worker.GetData().WordpressData)
}

// getWordpressPostbyIDHandler retrieves a WordPress post by ID
// @Summary Get WordPress post by ID
// @Description Get a specific WordPress post by its ID
// @Accept json
// @Produce json
// @Param id path string true "WordPress post ID"
// @Success 200 {object} object
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Post not found"
// @Router /wordpress/{id} [get]
func getWordpressPostbyIDHandler(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil || pid <= 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	} else if pid > len(worker.GetData().WordpressData) {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		wpData := worker.GetData().WordpressData
		c.JSON(http.StatusOK, wpData[len(wpData)-pid])
	}
}

// getSocials retrieves all social media links
// @Summary Get socials
// @Description Get all social media links
// @Accept json
// @Produce json
// @Success 200 {array} model.Social
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /socials [get]
func getSocials(c *gin.Context) {
	getResourcesBaseHandler[model.Social](socialsCollection, bson.M{"_id": -1})(c)
}
