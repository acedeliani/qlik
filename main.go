package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/items/:customerId", listItemsByCustomer)
	router.POST("/items/summary", listSummary)

	router.Run()
}

func listItemsByCustomer(c *gin.Context) {
	var orders Orders

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the items that belong to the customer
	// Assumption: the same customer may occur several times in the order collection
	customerId := c.Param("customerId")
	var customerItems []CustomerItem

	for _, o := range orders.Orders {
		if o.CustomerId == customerId {
			// Create and append a CustomerItem from an Item
			for _, i := range o.Items {
				customerItems = append(customerItems, CustomerItem{o.CustomerId, i.ItemId, i.CostEur})
			}
		}
	}

	c.IndentedJSON(http.StatusOK, &CustomerItemResponse{customerItems})
}

func listSummary(c *gin.Context) {
	fmt.Println("Inside listSummary")
}
