package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/product-management-system/api"
	"github.com/yourusername/product-management-system/cache"
	"github.com/yourusername/product-management-system/service"
	"github.com/yourusername/product-management-system/storage"
	"github.com/yourusername/product-management-system/queue"
	"github.com/yourusername/product-management-system/logs"
)

func main() {
	// Initialize logger
	logger := logs.InitLogger()

	// Initialize database connection
	db, err := storage.InitDB()
	if err != nil {
		logger.Fatal("failed to connect to database", err)
		os.Exit(1)
	}

	// Initialize Redis
	redisClient := cache.InitRedis()

	// Initialize message queue (RabbitMQ/Kafka)
	messageQueue, err := queue.InitQueue()
	if err != nil {
		logger.Fatal("failed to connect to message queue", err)
		os.Exit(1)
	}

	// Initialize product service
	productService := service.NewProductService(db, redisClient, messageQueue)

	// Create Gin router
	router := gin.Default()

	// Setup routes
	api.SetupRoutes(router, productService)

	// Start server
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("failed to start server", err)
	}
}
