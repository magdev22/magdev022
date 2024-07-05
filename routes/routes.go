package routes

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProductByID)
	router.DELETE("/products/:id", handlers.DeleteProductByID)
	router.POST("/updateProducts/:id", handlers.UpdateProductByID)
	router.POST("/products", handlers.PostProducts)
	router.Run("localhost:8080")
}
