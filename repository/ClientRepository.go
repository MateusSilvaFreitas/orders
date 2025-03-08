package repository

import (
	"database/sql"
	"fmt"

	"github.com/mateussilvafreitas/orders/database"
	model "github.com/mateussilvafreitas/orders/models"
)

func SaveClient(client model.Client) (int64, error){
	result, err := database.DB.Exec("INSERT INTO clients (name, email) VALUES (?, ? ) ", client.Name, client.Email)
	if err != nil {
		return 0, fmt.Errorf("saveClient: %v", err)
	}

	id, err := result.LastInsertId()
	
	if err != nil {
		return 0, fmt.Errorf("saveClient: %v", err)
	}

	return id, nil
}

func FindAllClients() ([]model.Client, error){
	var clients []model.Client

	rows, err := database.DB.Query("SELECT * FROM clients")

	if err != nil {
		return nil, fmt.Errorf("findAllClients: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var c model.Client

		if err := rows.Scan(&c.ID, &c.Name, &c.Email); err != nil {
			return nil, fmt.Errorf("findAllClients: %v", err)
		}

		clients = append(clients, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findAllClients: %v", err)
	}

	return clients, nil
}

func FindClientById(id int64) (model.Client, error){
	var client model.Client

	row := database.DB.QueryRow("SELECT id, name, email from clients where id=?", id)

	if err := row.Scan(&client.ID, &client.Name, &client.Email); err != nil {
		if err == sql.ErrNoRows {
			return client, fmt.Errorf("findClientById %d: no such client", id)
		}

		return client, fmt.Errorf("findClientById %d: %v", id, err)
	}

	return client, nil
}