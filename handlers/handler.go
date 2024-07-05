package handlers

import (
	model "api/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Products)
}

func PostProducts(c *gin.Context) {
	var NewProduct model.Product
	if err := c.BindJSON(&NewProduct); err != nil {
		return
	}
	model.Products = append(model.Products, NewProduct)
	c.IndentedJSON(http.StatusCreated, NewProduct)
}
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range model.Products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
func DeleteProductByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range model.Products {
		if a.ID == id {
			model.Products = append(model.Products[:i], model.Products[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
func UpdateProductByID(c *gin.Context) {
	var NewProduct model.Product
	id := c.Param("id")
	for i, a := range model.Products {
		if a.ID == id {
			c.BindJSON(&NewProduct)
			model.Products[i] = NewProduct
			c.IndentedJSON(http.StatusOK, NewProduct)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
