package model

type OrderProduct struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	OrderID   int64 `json:"order_id"`
	Quantity  int64 `json:"quantity"`
}