package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	model "github.com/mateussilvafreitas/orders/models"
	"github.com/mateussilvafreitas/orders/repository"
)

func PostSaveProduct(c *gin.Context){
	var product model.Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	id, err := repository.SaveProduct(product)

	if err != nil {
		c.IndentedJSON(500, gin.H{"message": fmt.Sprintf("Error saving product: %v", err)})
		return
	}

	product.ID = id
	c.IndentedJSON(201, product)
}

func FindAllProducts(c *gin.Context) {
	products, err := repository.FindAllProducts()

	if err != nil {
		c.IndentedJSON(500, gin.H{"message": fmt.Sprintf("Error finding all products: %v", err)})
		return
	}

	c.IndentedJSON(200, products)
}
