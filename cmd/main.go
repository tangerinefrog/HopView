package main

import (
	"log"
	"tangerinefrog/HopView/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env file: %v", err)
	}

	r := gin.Default()
	handlers.SetupRoutes(r)

	port := "8080"
	log.Printf("Server is starting on port :%s...\n", port)
	err = r.Run(":" + port)
	log.Printf("Error starting server: %v\n", err)
}
