package handlers

//product

import (
	model "api/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBills(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Bills)
}

func PostBills(c *gin.Context) {
	var NewBills model.Bill
	if err := c.BindJSON(&NewBills); err != nil {
		return
	}
	model.Bills = append(model.Bills, NewBills)
	c.IndentedJSON(http.StatusCreated, NewBills)
}
func GetBillByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range model.Bills {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
}
func DeleteBillByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range model.Bills {
		if a.ID == id {
			model.Bills = append(model.Bills[:i], model.Bills[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
}
func UpdateBillByID(c *gin.Context) {
	var NewBill model.Bill
	id := c.Param("id")
	for i, a := range model.Bills {
		if a.ID == id {
			c.BindJSON(&NewBill)
			model.Bills[i] = NewBill
			c.IndentedJSON(http.StatusOK, NewBill)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
}
