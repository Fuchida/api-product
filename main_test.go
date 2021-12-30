package main

import (
	"api_product/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductsRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	expected := 2
	var productList []domain.Product

	json.Unmarshal(w.Body.Bytes(), &productList)
	received := len(productList)
	
	testLog := "Testing /products route"
	testError := fmt.Sprintf("Listing products failed verification expected:%d but recieved: %d", expected, received)
	if expected != received {
		t.Logf(testLog)
		t.Errorf(testError)
	}
}
