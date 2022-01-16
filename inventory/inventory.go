package inventory

import (
	"api_product/domain"
	"errors"
	"fmt"
	"reflect"
)

var products = []domain.Product{}

// create initial stock for inventory
func Bootstrap() []domain.Product {
     products = []domain.Product{
        {ID: "21", Name: "Nintendo Switch Lite", Price: 399.90, Quantity: 2000},
        {ID: "18", Name: "Office Desk", Price: 200.00, Quantity: 5000},
    }
    return products
}

/* check if a products exists in the  inventory of products
*/ 		
func Exists(product domain.Product) bool {
/*
TODO: update to return bool for existance and index of product
    will enable other functions like deleteProduct to not loop twice
*/
    for _, item := range List() {
        if reflect.DeepEqual(item, product) {
            return true
        }
    }
    return false
}
// provided an id get the first available item with same ID
func GetInvetoryByID(id string) (domain.Product, error) {
    for _, item := range List() {
        if (item.ID == id) {
            return item, nil
        }
    }

    message := fmt.Sprintf("product with id %s not found", id)
    return domain.Product{}, errors.New(message)
}

// list the current available inventory
func List() []domain.Product {
    return products
}

// adds a new product to the inventory
func Add(product domain.Product) domain.Product {
    products = append(products, product	)
    
    return product
}

// updates an exsting inventory item
func Update(update domain.Product) domain.Product {
    // TODO: update func name to UpdateByID
    for index, product := range List() {
        if update.ID == product.ID {
            products[index] = update

            return update
        }
    }
    // MARK: Do we return an error when we fail to update ? 
    // 		 currently we just return an empty zero value Product
    return domain.Product{}
}

// removes the item from inventory
func Remove(index int) int {
    products = append(products[:index], products[index+1:]...)
    return index
}

// removes all items from inventory
func removeAll(){
    // MARK: RemoveAll is here as clean up for my tests but not used
    //       in other places. Is there a better way ?
    products = []domain.Product{}
}