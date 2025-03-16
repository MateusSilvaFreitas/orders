package repository

import (
	"database/sql"
	"fmt"

	"github.com/mateussilvafreitas/orders/database"
	model "github.com/mateussilvafreitas/orders/models"
)

func SaveOrder(order model.Order) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO orders (date_order, total_value, client_id) VALUES (?, ?, ?)", order.DateOrder, order.Total, order.ClientID)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("saveOrder: %v", err)
	}

	return id, nil
}


func SaveOrderProduct(orderProduct model.OrderProduct) (int64, error){
	result, err := database.DB.Exec("INSERT INTO order_product(product_id, order_id, quantity, unitary_price, total_price) VALUES(?,?,?,?,?)", orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity, orderProduct.UnitaryPrice, orderProduct.TotalPrice)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("saveOrderProduct: %v", err)
	}

	return id, nil
}

func FindOrderById(id int64) (model.Order, error) {
	var order model.Order

	row := database.DB.QueryRow("SELECT id, date_order, total_value, client_id FROM orders WHERE id=?", id)

	if err := row.Scan(&order.ID, &order.DateOrder, &order.Total, &order.ClientID); err != nil {
		if err == sql.ErrNoRows {
			return order, fmt.Errorf("findOrderById %d: no such order", id)
		}
		return order, fmt.Errorf("findOrderById %d: %v", id, err)
	}

	return order, nil
}

func FindProductsFromOrder(orderId int64) ([]model.OrderProduct, error) {
	var orderProducts []model.OrderProduct

	rows, err := database.DB.Query("SELECT id, product_id, order_id, quantity, unitary_price, total_price FROM order_product WHERE order_id=?", orderId)

	if err != nil {
		return nil, fmt.Errorf("findProductsFromOrder: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var op model.OrderProduct

		if err := rows.Scan(&op.ID, &op.ProductID, &op.OrderID, &op.Quantity, &op.UnitaryPrice, &op.TotalPrice); err != nil {
			return nil, fmt.Errorf("findProductsFromOrder: %v", err)
		}
		orderProducts = append(orderProducts, op)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findProductsFromOrder: %v", err)
	}
	
	return orderProducts, nil
}