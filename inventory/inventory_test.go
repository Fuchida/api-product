package inventory

import (
	"api_product/domain"
	"fmt"
	"testing"
)

func TestBoostrap(t *testing.T){
    expected := 2

    products = Bootstrap()
    received := len(products)
    defer removeAll()
    
    if received != expected {
        t.Logf("Testing boostaping the inventory")
        t.Errorf("Boostrap verification failed expected: %d but got: %d",
                expected, received)
    }
}

func TestExists(t *testing.T){
    _ = Bootstrap()
    defer removeAll()
    
    // validating existance
    nintendoSwitch := domain.Product{ID: "21", Name: "Nintendo Switch Lite", Price: 399.90, Quantity: 2000}
    expected := true
    recieved := Exists(nintendoSwitch)

    if expected != recieved {
        t.Logf("Testing existance of %s", nintendoSwitch.Name)
        t.Errorf("Existance verification failed expected: %t but recieved: %t", expected, recieved)
    }
    // validating non-existace
    sonyPlaystation := domain.Product{ID: "22", Name: "Sony Playstation", Price: 399.90, Quantity: 2000}
    expected = false
    recieved = Exists(sonyPlaystation)

    if expected != recieved {
        t.Logf("Testing existance of %s", sonyPlaystation.Name)
        t.Errorf("Existance verification failed expected: %t but recieved: %t", expected, recieved)
    }

}

func TestList(t *testing.T){
    expected := 0
    received := len(List())
    
    if expected != received {
        t.Logf("Testing listing of inventory items")
        t.Errorf("Listing verification failed expected: %d but recieved: %d", expected, received)
    }
    
    _ = Bootstrap()
    defer removeAll()
    expected = 2
    received = len(List())
    
    testLog := "Testing listing of inventory items"
    testError := fmt.Sprintf("Listing verification failed expected: %d but recieved: %d", expected, received)
    
    if expected != received {
        t.Logf(testLog)
        t.Errorf(testError)
    }
}

func TestAdd(t *testing.T){
    expected := true
    
    _ = Bootstrap()
    defer removeAll()
    
    segaGenesis := domain.Product{
        ID: "55",
        Name: "Sega Genesis",
        Price: 299.00,
        Quantity: 56,
        
    }
    
    _ = Add(segaGenesis)
    received := Exists(segaGenesis)
    
    testLog := "Testing adding an item to the inventory"
    testError := fmt.Sprintf("Inventory addition verification failed expected: %t but recieved: %t", expected, received)
    
    if expected != received {
        t.Logf(testLog)
        t.Errorf(testError)
    }
}

func TestUpdate(t *testing.T) {
    _ = Bootstrap()
    defer removeAll()
    
    expected := true
    
    nintendoSwitchPro := domain.Product{ID: "21", Name: "Nintendo Switch Pro", Price: 399.90, Quantity: 3000}
    
    _ = Update(nintendoSwitchPro)
    
    received := Exists(nintendoSwitchPro)
    
    testLog := "Testing updating an inventory item"
    testError := fmt.Sprintf("Inventory upate verification failed expected: %t but recieved: %t", expected, received)
    
    if expected != received {
        t.Logf(testLog)
        t.Errorf(testError)
    }
    
}

func TestRemove(t *testing.T){
    _ = Bootstrap()
    defer removeAll()
    
    expected := false
    SegaGameGear := domain.Product{ID: "33", Name: "Sega Game Gear", Price: 99.90, Quantity: 6000}
    
    _ = Add(SegaGameGear)
    for index, item := range List() {
        if item.ID == SegaGameGear.ID{
            Remove(index)
        }
    }
    received := Exists(SegaGameGear)
    
        
    testLog := "Testing removing an inventory item"
    testError := fmt.Sprintf("Inventory removal verification failed expected: %t but recieved: %t", expected, received)
    
    if expected != received {
        t.Logf(testLog)
        t.Errorf(testError)
    }
}

func TestRemoveAll(t *testing.T){
    _ = Bootstrap()
    expected := 0
    
    removeAll()
    received :=  len(List())
    
    testLog := "Testing removing an inventory item"
    testError := fmt.Sprintf("Inventory removal verification failed expected: %d but recieved: %d", expected, received)
    
    if expected != received {
        t.Logf(testLog)
        t.Errorf(testError)
    }
    
}