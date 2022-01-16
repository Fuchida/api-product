package main

import (
	"api_product/domain"
	"api_product/inventory"
	"bytes"
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

func TestAddProductRoute(t *testing.T) {
	expected := true

	router := setupRouter()
	snes := domain.Product{ID: "756", Name: "Super Nintendo", Price: 399.90, Quantity: 4000}
	payload, _ := json.Marshal(snes)

	product := []byte(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.ServeHTTP(w, req)

	testLog := "Testing POST on /product route"
	testError := fmt.Sprintf("Adding product failed expected inventory ID: %s but inventory check failed with %t",
		snes.ID, inventory.Exists(snes))

	if inventory.Exists(snes) != expected {
		t.Logf(testLog)
		t.Errorf(testError)
	}
}
