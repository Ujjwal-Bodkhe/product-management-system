package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/product-management-system/service"
)

// SetupRoutes sets up the product API routes
func SetupRoutes(router *gin.Engine, productService *service.ProductService) {
	router.POST("/products", createProductHandler(productService))
	router.GET("/products/:id", getProductHandler(productService))
	router.GET("/products", getProductsHandler(productService))
}
