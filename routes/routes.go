package routes

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/bills", handlers.GetBills)
	router.GET("/bills/:id", handlers.GetBillByID)
	router.DELETE("/bills/:id", handlers.DeleteBillByID)
	router.POST("/updateBills/:id", handlers.UpdateBillByID)
	router.POST("/bills", handlers.PostBills)
	router.Run("localhost:8080")
}
