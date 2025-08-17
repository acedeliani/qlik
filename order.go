package main

type Item struct {
	ItemId  string  `json:"itemId" binding:"required"`
	CostEur float32 `json:"costEur" binding:"required"`
}

type Order struct {
	CustomerId string `json:"customerId" binding:"required"`
	OrderId    string `json:"orderId" binding:"required"`
	Timestamp  string `json:"timestamp" binding:"required"`
	Items      []Item `json:"items" binding:"required,dive"`
}

type Orders struct {
	Orders []Order `json:"orders" binding:"required,dive"`
}

// Responses
type CustomerItem struct {
	CustomerId string  `json:"customerId"`
	ItemId     string  `json:"itemId"`
	CostEur    float32 `json:"costEur"`
}

// Response for /items/:customerId
type CustomerItemResponse struct {
	Items []CustomerItem `json:"items" binding:"required,dive"`
}
