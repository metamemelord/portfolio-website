package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/metamemelord/portfolio-website/worker"
)

func refreshData(c *gin.Context) {
	log.Println("Force refresh called")
	worker.RefreshData()
	respond(c, http.StatusAccepted, nil, nil)
}

func getGithubReposHandler(c *gin.Context) {
	log.Println("[INFO] Github data accessed")
	c.JSON(http.StatusOK, worker.GetData().GithubData)
}

func getWordpressPostsHandler(c *gin.Context) {
	log.Println("[INFO] Wordpress data accessed")
	c.JSON(http.StatusOK, worker.GetData().WordpressData)
}
