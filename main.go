package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mateussilvafreitas/orders/database"
	"github.com/mateussilvafreitas/orders/service"
)

func main() {
	database.InitDatabase()

	router := gin.Default()

	router.GET("/products", service.FindAllProducts)
	router.GET("/clients", service.GetFindAllClients)
	router.GET("/clients/:id", service.GetFindClientById)
	router.GET("/orders/:id", service.GetOrderById)

	router.POST("/clients", service.PostSaveClient)
	router.POST("/products", service.PostSaveProduct)
	router.POST("/orders", service.PostCreateOrder)


	router.Run("localhost:8080")

	fmt.Println("Application started!")
}