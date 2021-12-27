/*

Package main implements a HTTP web service api.
	. List existing products ✅
	. Lookup product by id ✅
	. Add new product  ✅
	. Reject duplicate product ✅
	. Delete product by id ✅
	. Update existing product
	. Add unit tests for above functionality
	. log happy paths and failure paths
	. Persist products to an sqlite database (create, update, delete)
	. send events to messaging system
		- new product added
		- product deleted
		- product updated
	. Add metrics to public endpoints
		- list
		- add
		- get
		- delete
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int32   `json:"quantity"`
}

// removes a product from  inventory
func removeProductByIndex(index int) {
	products = append(products[:index], products[index+1:]...)

}

// check if a products exists in the  inventory of products
// TODO: update to return bool for existance and index of product
// 		will enable other functions like deleteProduct to not loop twice
func productExists(p product, items []product) bool {
	for _, product := range items {
		if product.ID == p.ID {
			return true
		}
	}
	return false
}

var products = []product{
	{ID: "21", Name: "Nintendo Switch", Price: 399.90, Quantity: 2000},
	{ID: "18", Name: "Office Desk", Price: 200.00, Quantity: 5000},
}

func addProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	if productExists(newProduct, products) {
		c.IndentedJSON(http.StatusConflict, "product already exists")
		return

	} else {
		products = append(products, newProduct)
		c.IndentedJSON(http.StatusAccepted, newProduct)
	}
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.IndentedJSON(http.StatusOK, product)
			return
		}
	}
	message := fmt.Sprintf("Product id: %s, was not found", id)
	c.IndentedJSON(http.StatusNotFound, message)
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	for index, product := range products {
		if product.ID == id {
			removeProductByIndex(index)
			message := fmt.Sprintf("Product id: %s, has been deleted", product.ID)

			c.IndentedJSON(http.StatusAccepted, message)
			return
		}
	}
	message := fmt.Sprintf("Product id: %s, was not found ", id)
	c.IndentedJSON(http.StatusNotFound, message)
}

func updateProduct(c *gin.Context) {
	var patch product

	if err := c.BindJSON(&patch); err != nil {
		return
	}

	if productExists(patch, products) {
		for index, product := range products {
			if patch.ID == product.ID {
				products[index] = patch

				// TODO: can we log the full representation of the product ?
				message := fmt.Sprintf("Product: %s, has been updated", product.ID)

				c.IndentedJSON(http.StatusAccepted, message)
				return
			}
		}
	} else {
		message := fmt.Sprintf("Product id: %s, was not found ", patch.ID)
		c.IndentedJSON(http.StatusNotFound, message)
	}
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/product/:id", getProductByID)
	router.POST("/product", addProduct)
	router.PATCH("/product/", updateProduct)
	router.DELETE("product/:id", deleteProduct)
	router.Run("localhost:7000")
}
