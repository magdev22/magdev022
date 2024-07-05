package main

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/products", handlers.GetProducts)
	// 	router.GET("/products/:id", getProductByID)
	// 	router.DELETE("/products/:id", deleteProductByID)
	// 	router.POST("/updateProducts/:id", updateProductByID)
	// 	router.POST("/products", postProducts)
	// 	router.Run("localhost:8080")
}
