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

func TestPatchProductRoute(t *testing.T) {
	expected := true

	router := setupRouter()
	xbox := domain.Product{ID: "929", Name: "Xbox", Price: 200.00, Quantity: 7000}

	_ = inventory.Add(xbox)

	xbox.Price = 300.00
	payload, _ := json.Marshal(xbox)
	patchedProduct := []byte(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/product", bytes.NewBuffer(patchedProduct))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.ServeHTTP(w, req)

	testLog := "Testing PATCH on /product route"
	testError := fmt.Sprintf("Adding product failed expected inventory price: %f but inventory check failed with %t",
		xbox.Price, inventory.Exists(xbox))

	if inventory.Exists(xbox) != expected {
		t.Logf(testLog)
		t.Errorf(testError)
	}
}

func TestDeleteProductRoute(t *testing.T) {
	expected := false

	PS5 := domain.Product{ID: "800", Name: "PS5", Price: 400.00, Quantity: 17000}
	deletePath := fmt.Sprintf("/product/%s", PS5.ID)

	router := setupRouter()
	_ = inventory.Add(PS5)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.ServeHTTP(w, req)

	testLog := "Testing Delete on /product route with id"
	testError := fmt.Sprintf("Deleting product failed for %s inventory check returned  %t",
		PS5.Name, inventory.Exists(PS5))

	if inventory.Exists(PS5) != expected {
		t.Logf(testLog)
		t.Errorf(testError)
	}

}
