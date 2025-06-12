package handlers

import (
	"net/http"
	"tangerinefrog/HopView/internal/network"

	"github.com/gin-gonic/gin"
)

func tracerouteHandler(c *gin.Context) {
	url := c.Query("url")
	nodes, err := network.TraceRoute(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not get route values from provided URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"endpoints": nodes})
}
