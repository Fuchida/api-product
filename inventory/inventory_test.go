package inventory

import (
	"api_product/domain"
	"testing"
)

var _ []domain.Product = Bootstrap()

func Test_Exists(t *testing.T){
	// validating existance
	nintendoSwitch := domain.Product{ID: "21", Name: "Nintendo Switch Lite", Price: 399.90, Quantity: 2000}
	want := true
	got := Exists(nintendoSwitch)
	
	if want != got {
		t.Logf("Testing existance of %s", nintendoSwitch.Name)
		t.Errorf("Existance verification failed want: %t but got: %t", want, got)
	}
	// validating non-existace
	sonyPlaystation := domain.Product{ID: "22", Name: "Sony Playstation", Price: 399.90, Quantity: 2000}
	want = false
	got = Exists(sonyPlaystation)
	
	if want != got {
		t.Logf("Testing existance of %s", sonyPlaystation.Name)
		t.Errorf("Existance verification failed want: %t but got: %t", want, got)
	}
	
}