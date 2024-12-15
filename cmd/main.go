package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Ujjwal-Bodkhe/product-management-system/cache"
	"github.com/Ujjwal-Bodkhe/product-management-system/handlers"
	"github.com/Ujjwal-Bodkhe/product-management-system/queue"
	"github.com/Ujjwal-Bodkhe/product-management-system/service"
	"github.com/Ujjwal-Bodkhe/product-management-system/storage"
)

func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize components
	db := storage.NewDB()
	redisClient := cache.InitRedis()
	messageQueue := queue.NewMessageQueue()

	// Create product service
	productService := service.NewProductService(db, redisClient, messageQueue)

	// Setup router and routes
	router := mux.NewRouter()
	handlers.SetupProductRoutes(router, productService)

	// Start server
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadEnv() error {
	// Load environment variables from .env file
	return nil // Placeholder, you can use a package like "github.com/joho/godotenv" to load .env file
}
