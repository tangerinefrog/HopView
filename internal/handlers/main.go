package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
