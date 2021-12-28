package inventory

import "api_product/domain"

// create initial stock for inventory
func Bootstrap() []domain.Product {
	var products = []domain.Product{
		{ID: "21", Name: "Nintendo Switch Lite", Price: 399.90, Quantity: 2000},
		{ID: "18", Name: "Office Desk", Price: 200.00, Quantity: 5000},
	}
	return products
}

// check if a products exists in the  inventory of products
// TODO: update to return bool for existance and index of product
// 		will enable other functions like deleteProduct to not loop twice
func ProductExists(p domain.Product, items []domain.Product) bool {
	for _, product := range items {
		if product.ID == p.ID {
			return true
		}
	}
	return false
}
