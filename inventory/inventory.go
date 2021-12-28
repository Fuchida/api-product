package inventory

import (
	"api_product/domain"
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

// check if a products exists in the  inventory of products
// TODO: update to return bool for existance and index of product
// 		will enable other functions like deleteProduct to not loop twice
func Exists(p domain.Product) bool {
	for _, product := range List() {
		if product.ID == p.ID {
			return true
		}
	}
	return false
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