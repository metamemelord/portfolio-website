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
	respond(c, http.StatusOK, worker.GetData().GithubData, nil)
}

func getWordpressPostsHandler(c *gin.Context) {
	respond(c, http.StatusOK, worker.GetData().WordpressData, nil)
}
