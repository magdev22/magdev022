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
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func deleteProductByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range products {
		if a.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)

			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateProductByID(c *gin.Context) {
	id := c.Param("id")

	for i := range products {
		if products[i].ID == id {
			c.BindJSON(&products[i])
			c.IndentedJSON(http.StatusOK, products[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.DELETE("/products/:id", deleteProductByID)
	router.PUT("/products/:id", updateProductByID)
	router.POST("/products", postProducts)
	router.Run("localhost:8080")
}
