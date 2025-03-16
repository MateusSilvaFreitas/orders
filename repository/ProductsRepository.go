package repository

import (
	"database/sql"
	"fmt"

	"github.com/mateussilvafreitas/orders/database"
	model "github.com/mateussilvafreitas/orders/models"
)

func SaveProduct(product model.Product) (int64, error){
	result, err := database.DB.Exec("INSERT INTO products (name,price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		return 0, fmt.Errorf("saveProduct: %v", err)
	}

	id, err := result.LastInsertId();

	if err != nil {
		return 0, fmt.Errorf("saveProduct: %v", err)
	}

	return id, nil
}

func FindAllProducts() ([]model.Product, error){
	var products []model.Product

	rows, err := database.DB.Query("SELECT * FROM products")

	if err != nil {
		return nil, fmt.Errorf("findAllProducts: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, fmt.Errorf("findAllProducts: %v", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findAllProducts: %v", err)
	}

	return products, nil
}

func FindProductById(id int64) (model.Product, error){
	var product model.Product
	
	row := database.DB.QueryRow("SELECT id, name, price from products where id=?", id)

	if err:= row.Scan(&product.ID, &product.Name, &product.Price); err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("findProductById %d: no such product", id)
		}
		return product, fmt.Errorf("findProductById %d: %v", id, err)
	}

	return product, nil

}