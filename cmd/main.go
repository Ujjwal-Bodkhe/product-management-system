package main

import (
	"log"
	"github.com/Ujjwal-Bodkhe/product-management-system/service"
	"github.com/Ujjwal-Bodkhe/product-management-system/storage"
	"github.com/Ujjwal-Bodkhe/product-management-system/cache"
	"github.com/Ujjwal-Bodkhe/product-management-system/queue"
	"github.com/Ujjwal-Bodkhe/product-management-system/api"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Initialize components
	db := storage.InitDB()
	redisClient := cache.InitRedis()
	messageQueue := queue.InitMessageQueue()

	// Initialize ProductService with dependencies
	productService := service.NewProductService(db, redisClient, messageQueue)

	// Initialize the HTTP handler with the ProductService
	productHandler := api.NewProductHandler(productService)

	// Setup routes
	router := mux.NewRouter()
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/products/user/{userID}", productHandler.GetProductsByUser).Methods("GET")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
