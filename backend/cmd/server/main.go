package main

import (
	"backend/api/handlers"
	"backend/internal/config"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Pack Calculator API
// @version 1.0
// @description API for calculating and managing pack sizes.
// @host localhost:8080
// @BasePath /
// @schemes http
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

	// Swagger endpoint
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Define routes
	r.POST("/calculate", handlers.CalculateHandler(config))
	r.POST("/packs", handlers.UpdatePackSizesHandler(config))

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
