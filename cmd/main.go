package main

import (
	"log"
	"tangerinefrog/HopView/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handlers.SetupRoutes(r)

	port := "8080"
	log.Printf("Server is starting on port :%s...\n", port)
	err := r.Run(":" + port)
	log.Printf("Error starting server: %v\n", err)
}
