package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLFiles("web/index.html")
	r.Static("/static", "./web/static")

	r.GET("/ping", pingHandler)
	r.GET("/", mainHandler)
	r.GET("api/traceroute", tracerouteHandler)
}
