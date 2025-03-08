package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	model "github.com/mateussilvafreitas/orders/models"
	repository "github.com/mateussilvafreitas/orders/repository"
)




func PostSaveClient(c *gin.Context){
	var client model.Client

	if err := c.BindJSON(&client); err != nil {
		return
	}

	id, err := repository.SaveClient(client)

	if err != nil {
		c.IndentedJSON(500, gin.H{"message": fmt.Sprintf("Error saving client:  %v", err)})
		return
	}

	client.ID = id
	c.IndentedJSON(201, client)	
}

func GetFindAllClients(c *gin.Context){
	clients, err := repository.FindAllClients()

	if err != nil {
		c.IndentedJSON(500, gin.H{"message": fmt.Sprintf("Error finding all clients: %v", err)})
		return
	}

	c.IndentedJSON(200, clients)
}

func GetFindClientById(c *gin.Context){
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 0)
	
	if(err != nil) {
		c.IndentedJSON(400, gin.H{"message": "Invalid id"})
		return
	}

	client, err := repository.FindClientById(id)

	if(err != nil) {
		c.IndentedJSON(404, gin.H{"message": fmt.Sprintf("Error finding client by id: %v", err)})
		return
	}

	c.IndentedJSON(200, client)
}