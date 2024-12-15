package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Ujjwal-Bodkhe/product-management-system/service"
	"github.com/Ujjwal-Bodkhe/product-management-system/models"
	"github.com/gorilla/mux"
)

func SetupProductRoutes(router *mux.Router, productService *service.ProductService) {
	router.HandleFunc("/products", productService.CreateProductHandler).Methods("POST")
	router.HandleFunc("/products/{id}", productService.GetProductByIDHandler).Methods("GET")
	router.HandleFunc("/products", productService.GetProductsByUserHandler).Methods("GET")
}

func (s *ProductService) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call the CreateProduct method
	err := s.CreateProduct(&product)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *ProductService) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, err := s.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (s *ProductService) GetProductsByUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	products, err := s.GetProductsByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}
