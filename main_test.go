package main

import (
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

	expected := 200
	received := w.Code

	testLog := "Testing /products route"
	testError := fmt.Sprintf("Listing products failed verification expected:%d but recieved: %d", expected, received)

	if expected != received {
		t.Logf(testLog)
		t.Errorf(testError)
	}
}
