package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	model "github.com/mateussilvafreitas/orders/models"
	repository "github.com/mateussilvafreitas/orders/repository"
	"github.com/mateussilvafreitas/orders/utils"
)




func PostSaveClient(c *gin.Context){
	var client model.Client

	if err := c.BindJSON(&client); err != nil {
		return
	}

	id, err := repository.SaveClient(client)

	if err != nil {
		utils.HandleError(c, 500, "Error saving client", err)
		return
	}

	client.ID = id
	c.IndentedJSON(201, client)	
}

func GetFindAllClients(c *gin.Context){
	clients, err := repository.FindAllClients()

	if err != nil {
		utils.HandleError(c, 500, "Error finding all clients", err)
		return
	}

	c.IndentedJSON(200, clients)
}

func GetFindClientById(c *gin.Context){
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 0)
	
	if(err != nil) {
		utils.HandleError(c, 400, "Invalid id", err)
		return
	}

	client, err := repository.FindClientById(id)

	if(err != nil) {
		utils.HandleError(c, 404, "Error finding client by id", err)
		return
	}

	c.IndentedJSON(200, client)
}