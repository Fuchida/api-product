/*

Package main implements a HTTP web service api.
    . List existing products ✅
    . Lookup product by id ✅
    . Reject duplicate product ✅
    . Add new product  ✅
    . Delete product by id ✅
    . Update existing product ✅
    . Move services and domains to own packages ✅
    . Add unit tests for above functionality
		. packages ✅
		. api routes
    . log happy paths and failure paths
    . Persist products to an sqlite database (create, update, delete)
    . send events to messaging system
        - new product added
        - product deleted
        - product updated
    . Add promethius metrics for public endpoints
        - list
        - add
        - get
        - delete
*/
package main

import (
	"api_product/domain"
	"api_product/inventory"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ []domain.Product = inventory.Bootstrap()

func addProduct(c *gin.Context) {
	var newProduct domain.Product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	if inventory.Exists(newProduct) {
		c.IndentedJSON(http.StatusConflict, "product already exists")
		return

	} else {
		_ = inventory.Add(newProduct)
		c.IndentedJSON(http.StatusAccepted, newProduct)
	}
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range inventory.List() {
		if product.ID == id {
			c.IndentedJSON(http.StatusOK, product)
			return
		}
	}
	message := fmt.Sprintf("Product id: %s, was not found", id)
	c.IndentedJSON(http.StatusNotFound, message)
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, inventory.List())
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	for index, product := range inventory.List() {
		if product.ID == id {
			_ = inventory.Remove(index)
			message := fmt.Sprintf("Product id: %s, has been deleted", product.ID)

			c.IndentedJSON(http.StatusAccepted, message)
			return
		}
	}
	message := fmt.Sprintf("Product id: %s, was not found ", id)
	c.IndentedJSON(http.StatusNotFound, message)
}

func updateProduct(c *gin.Context) {
	var patch domain.Product

	if err := c.BindJSON(&patch); err != nil {
		return
	}

	if inventory.Exists(patch) {
		// MARK: can we log the full representation of the product ?
		product := inventory.Update(patch)
		message := fmt.Sprintf("Product: %s, has been updated", product.ID)

		c.IndentedJSON(http.StatusAccepted, message)

	} else {
		message := fmt.Sprintf("Product id: %s, was not found ", patch.ID)
		c.IndentedJSON(http.StatusNotFound, message)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/products", getProducts)
	router.GET("/product/:id", getProductByID)
	router.POST("/product", addProduct)
	router.PATCH("/product/", updateProduct)
	router.DELETE("product/:id", deleteProduct)
	return router
}

func main() {
	router := setupRouter()
	router.Run(":7000")
}
