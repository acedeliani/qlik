package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Suppress debug printouts during tests
	gin.SetMode(gin.TestMode)

	code := m.Run()

	os.Exit(code)
}

func readTestData(t *testing.T, name string) []byte {
	t.Helper()

	content, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Errorf("Could not read '%v'", name)
	}

	return content
}

func TestListItemsByCustomer(t *testing.T) {
	router := setupRouter()
	data := readTestData(t, "order.json")
	customerId := "01"

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/items/%s", customerId), bytes.NewReader(data))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Verify body
	expected := CustomerItemResponse{
		Items: []*CustomerItem{
			{CustomerId: customerId, ItemId: "20201", CostEur: 2},
			{CustomerId: customerId, ItemId: "20202", CostEur: 1.5},
			{CustomerId: customerId, ItemId: "50242", CostEur: 15},
		},
	}

	var responseBody CustomerItemResponse
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	assert.Equal(t, expected, responseBody)
}

func TestListItemsByCustomerCustomerNotFound(t *testing.T) {
	router := setupRouter()
	data := readTestData(t, "order.json")
	customerId := "05" // Invalid customer ID

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/items/%s", customerId), bytes.NewReader(data))
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	expectedJson := fmt.Sprintf("{\"error\":\"Could not find customer with ID '%s'\"}", customerId)
	assert.Equal(t, expectedJson, w.Body.String())
}

func TestListItemsByCustomerInvalidJson(t *testing.T) {
	router := setupRouter()
	data := readTestData(t, "order-missing-timestamp.json")
	customerId := "01"

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/items/%s", customerId), bytes.NewReader(data))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	expected := "{\"error\":\"Key: 'Orders.Orders[0].Timestamp' Error:Field validation for 'Timestamp' failed on the 'required' tag\"}"
	assert.Equal(t, expected, w.Body.String())
}

func TestListSummary(t *testing.T) {
	router := setupRouter()
	data := readTestData(t, "order.json")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/items/summary", bytes.NewReader(data))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	expected := ItemSummaryResponse{
		Items: []*CustomerSummary{
			{CustomerId: "01", NbrOfPurchasedItems: 3, TotalAmountEur: 18.5},
			{CustomerId: "02", NbrOfPurchasedItems: 1, TotalAmountEur: 5},
		},
	}

	var responseBody ItemSummaryResponse
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	assert.Equal(t, expected, responseBody)
}

func TestListSummaryInvalidJson(t *testing.T) {
	router := setupRouter()
	data := readTestData(t, "order-missing-timestamp.json")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/items/summary", bytes.NewReader(data))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	expected := "{\"error\":\"Key: 'Orders.Orders[0].Timestamp' Error:Field validation for 'Timestamp' failed on the 'required' tag\"}"
	assert.Equal(t, expected, w.Body.String())
}
