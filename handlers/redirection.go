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
