package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/product-management-system/service"
)

// createProductHandler handles POST requests to create a new product
func createProductHandler(productService *service.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product service.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := productService.CreateProduct(&product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
	}
}

// getProductHandler handles GET requests to retrieve a product by ID
func getProductHandler(productService *service.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product, err := productService.GetProductByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// getProductsHandler handles GET requests to retrieve all products for a user
func getProductsHandler(productService *service.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.DefaultQuery("user_id", "")
		products, err := productService.GetProductsByUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}
