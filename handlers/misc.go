package handlers

import (
	"log"
	"net/http"
	"strconv"

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
