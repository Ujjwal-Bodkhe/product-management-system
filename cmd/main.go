package main

import (
	"log"
	"net/http"
	"github.com/Ujjwal-Bodkhe/product-management-system/service"
	"github.com/Ujjwal-Bodkhe/product-management-system/handler"
	"github.com/Ujjwal-Bodkhe/product-management-system/storage"
	"github.com/Ujjwal-Bodkhe/product-management-system/cache"
	"github.com/Ujjwal-Bodkhe/product-management-system/queue"
	"github.com/Ujjwal-Bodkhe/product-management-system/routes"
)

func main() {
	// Initialize database connection, Redis client, and message queue
	db := storage.NewDB("postgres://user:pass@localhost/dbname")
	redisClient := cache.NewRedisClient("localhost:6379")
	messageQueue := queue.NewMessageQueue("localhost:5672")

	// Initialize services
	productService := service.NewProductService(db, redisClient, messageQueue)

	// Initialize handlers
	productHandler := handler.NewProductHandler(productService)

	// Set up routes
	router := routes.NewRouter(productHandler)

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
