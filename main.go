package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"product"`
	Price  float64 `json:"price"`
}

var products = []product{
	{ID: "1", Title: "kanfetka", Artist: "chiorio", Price: 52.52},
	{ID: "2", Title: "moloko", Artist: "nemoloko", Price: 77.77},
	{ID: "3", Title: "chokoladka", Artist: "alenka", Price: 66.66},
}

func main() {
	router := gin.Default()
	router.GET("/product", getProducts)

	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}
