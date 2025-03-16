package model

type OrderProduct struct {
	ID           int64   `json:"id"`
	ProductID    int64   `json:"productId"`
	OrderID      int64   `json:"orderId"`
	Quantity     int64   `json:"quantity"`
	UnitaryPrice float64 `json:"unitaryPrice"`
	TotalPrice   float64 `json:"totalPrice"`
}