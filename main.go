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

	router.GET("/clients", service.GetFindAllClients)
	router.GET("/clients/:id", service.GetFindClientById)
	router.POST("/clients", service.PostSaveClient)

	router.Run("localhost:8080")

	fmt.Println("Application started!")
}