package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/metamemelord/portfolio-website/model"
	"github.com/metamemelord/portfolio-website/pkg/core"
	"github.com/metamemelord/portfolio-website/pkg/worker"
)

// addRedirection adds a new URL redirection
// @Summary Add a new redirection
// @Description Create a new URL redirection (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param redirection body model.RedirectionItem true "Redirection data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 409 {object} map[string]interface{} "Route already exists"
// @Failure 422 {object} map[string]interface{} "Unprocessable entity"
// @Router /redirect [post]
func addRedirection(c *gin.Context) {
	redirectionItem := model.RedirectionItem{}
	if err := c.BindJSON(&redirectionItem); err != nil {
		log.Println(err)
		respond(c, http.StatusBadRequest, nil, err)
	} else {
		if redirectionItem.RoutingKey == core.EMPTY_STRING {
			respond(c, http.StatusBadRequest, nil, errors.New("Empty routing_key provided"))
			return
		}

		target, _, _ := worker.ResolveRedirectionItem(redirectionItem.RoutingKey, core.EMPTY_STRING, core.EMPTY_STRING)
		if target != core.EMPTY_STRING {
			respond(c, http.StatusConflict, nil, errors.New("This route already exists"))
			return
		}
		_, err := worker.AddRedirectionItem(c.Request.Context(), &redirectionItem)
		if err != nil {
			respond(c, http.StatusUnprocessableEntity, nil, err)
		} else {
			respond(c, http.StatusCreated, map[string]string{core.ROUTING_KEY: redirectionItem.RoutingKey, "_id": redirectionItem.ID.Hex()}, nil)
		}
	}
}

// getRedirectionItems retrieves all redirections
// @Summary Get redirections
// @Description Get all redirections or only active ones (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param all query string false "Include inactive redirections (true or 1)"
// @Success 200 {array} model.RedirectionItem
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 503 {object} map[string]interface{} "Service unavailable"
// @Router /redirect [get]
func getRedirectionItems(c *gin.Context) {
	allValues := strings.ToLower(c.Query("all"))
	filters := make([]model.RedirectionItemSearchFilter, 0)
	if !(allValues == "true" || allValues == "1") {
		filters = append(filters, model.RedirectionItemSearchFilter{Key: core.ACTIVE, Value: true})
	}

	redirectionItems, err := worker.GetRedirectionItems(c.Request.Context(), filters...)
	if err != nil {
		respond(c, http.StatusServiceUnavailable, nil, err)
	} else {
		respond(c, http.StatusOK, redirectionItems, nil)
	}
}

// getRedirectionItemByRoutingKey retrieves a redirection by routing key
// @Summary Get redirection by routing key
// @Description Get a specific redirection by its routing key (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param routing_key path string true "Routing key"
// @Param active query string false "Only active redirections (true or 1)"
// @Success 200 {array} model.RedirectionItem
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 503 {object} map[string]interface{} "Service unavailable"
// @Router /redirect/{routing_key} [get]
func getRedirectionItemByRoutingKey(c *gin.Context) {
	routingKey := c.Param(core.ROUTING_KEY)
	filters := []model.RedirectionItemSearchFilter{{Key: core.ROUTING_KEY, Value: routingKey}}

	activeOnly := c.Query("active")
	if activeOnly == "true" || activeOnly == "1" {
		filters = append(filters, model.RedirectionItemSearchFilter{Key: core.ACTIVE, Value: true})
	}

	redirectionItems, err := worker.GetRedirectionItems(c.Request.Context(), filters...)
	if err != nil {
		respond(c, http.StatusServiceUnavailable, nil, err)
	} else {
		respond(c, http.StatusOK, redirectionItems, nil)
	}
}

func resolveRedirection(c *gin.Context) {
	routingKey := c.Param(core.ROUTING_KEY)
	pathToForward := c.Param(core.PATH_LABEL)
	if target, code, err := worker.ResolveRedirectionItem(routingKey, pathToForward, c.Request.URL.RawQuery); err == nil {
		c.Redirect(code, target)
		c.Abort()
	} else {
		log.Println("Error while redirecting", err)
	}
}

// deleteRedirection deletes a redirection by routing key
// @Summary Delete redirection
// @Description Delete a redirection by its routing key (requires authentication)
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param routing_key path string true "Routing key"
// @Success 200
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 503 {object} map[string]interface{} "Service unavailable"
// @Router /redirect/{routing_key} [delete]
func deleteRedirection(c *gin.Context) {
	routingKey := c.Param(core.ROUTING_KEY)
	if err := worker.DeleteRedirectionItem(routingKey); err == nil {
		respond(c, http.StatusOK, nil, nil)
	} else {
		respond(c, http.StatusServiceUnavailable, nil, err)
	}
}

func SubdomainRedirectionMiddleware(c *gin.Context) {
	hostname := c.Request.Host
	if isSubdomainRequest(hostname) {
		code, uri := worker.GetSubdomainRedirection(hostname, c.Request.URL.RequestURI())
		c.Redirect(code, uri)
		c.Abort()
	} else {
		c.Next()
	}
}

func isSubdomainRequest(hostname string) bool {
	if subdomainPattern.Match([]byte(hostname)) {
		return !strings.HasPrefix("www.", hostname)
	}
	return false
}
