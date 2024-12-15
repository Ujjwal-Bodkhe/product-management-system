package routes

import (
	"github.com/gorilla/mux"
	"github.com/Ujjwal-Bodkhe/product-management-system/handler"
)

func NewRouter(productHandler *handler.ProductHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/products/user/{user_id}", productHandler.GetProductsByUser).Methods("GET")

	return router
}
