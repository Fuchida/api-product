package main

import (
	"api_product/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestGetProductByIDRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/product/21", nil)
	router.ServeHTTP(w, req)

	expectedProduct := domain.Product{ID: "21", Name: "Nintendo Switch Lite", Price: 399.90, Quantity: 2000}
	var receivedProduct domain.Product

	json.Unmarshal(w.Body.Bytes(), &receivedProduct)

	testLog := "Testing /product/:id route"
	testError := fmt.Sprintf("Listing product via id failed verification expected ID: %s but received ID: %s",
		expectedProduct.ID, receivedProduct.ID)

	if !reflect.DeepEqual(expectedProduct, receivedProduct) {
		t.Logf(testLog)
		t.Errorf(testError)
	}

}
