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
	var customerItems []*CustomerItem

	for _, o := range orders.Orders {
		if o.CustomerId == customerId {
			// Create and append a CustomerItem from an Item
			for _, i := range o.Items {
				customerItems = append(customerItems, &CustomerItem{o.CustomerId, i.ItemId, i.CostEur})
			}
		}
	}

	if len(customerItems) == 0 {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"error": fmt.Sprintf("Could not find customer with ID '%s'", customerId)},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, &CustomerItemResponse{customerItems})
}

func listSummary(c *gin.Context) {
	var orders Orders

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerMap := summaryMapFromOrders(orders)

	var items []*CustomerSummary
	for _, v := range *customerMap {
		items = append(items, v)
	}

	c.IndentedJSON(http.StatusOK, items)
}

func summaryMapFromOrders(orders Orders) *map[string]*CustomerSummary {
	customerMap := make(map[string]*CustomerSummary)

	for _, o := range orders.Orders {
		if _, ok := customerMap[o.CustomerId]; !ok {
			customerMap[o.CustomerId] = &CustomerSummary{CustomerId: o.CustomerId}
		}

		customerMap[o.CustomerId].NbrOfPurchasedItems += len(o.Items)

		for _, i := range o.Items {
			customerMap[o.CustomerId].TotalAmountEur += i.CostEur
		}
	}

	return &customerMap
}
