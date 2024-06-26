package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}

var products = []product{
	{ID: "1", Title: "kanfetka", Product: "chi orio", Price: 52.52},
	{ID: "2", Title: "moloko", Product: "ne moloko", Price: 77.77},
	{ID: "3", Title: "chokoladka", Product: "alenka", Price: 66.66},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/products", postProducts)
	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func postProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)

}
