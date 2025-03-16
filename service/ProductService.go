package service

import (
	"github.com/gin-gonic/gin"
	model "github.com/mateussilvafreitas/orders/models"
	"github.com/mateussilvafreitas/orders/repository"
	"github.com/mateussilvafreitas/orders/utils"
)

func PostSaveProduct(c *gin.Context){
	var product model.Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	id, err := repository.SaveProduct(product)

	if err != nil {
		utils.HandleError(c, 500, "Error saving product", err)		
		return
	}

	product.ID = id
	c.IndentedJSON(201, product)
}

func FindAllProducts(c *gin.Context) {
	products, err := repository.FindAllProducts()

	if err != nil {
		utils.HandleError(c, 500, "Error finding all products",  err)
		return
	}

	c.IndentedJSON(200, products)
}
