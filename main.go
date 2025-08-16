package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world!")

	router := gin.Default()

	router.POST("/items/:customerId", listItemsByCustomer)
	router.POST("/items/summary", listSummary)

	router.Run()
}

func listItemsByCustomer(c *gin.Context) {
	fmt.Println("Inside listItemsByCustomer")
}

func listSummary(c *gin.Context) {
	fmt.Println("Inside listSummary")
}
