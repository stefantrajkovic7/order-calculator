package main

import (
	"backend/api/handlers"
	"backend/internal/config"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("./internal/config/pack_sizes.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())

	// Define routes
	r.POST("/calculate", handlers.CalculateHandler(config))
	r.POST("/packs", handlers.UpdatePackSizesHandler(config))

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
