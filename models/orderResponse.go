package model

type OrderResponse struct {
	Id         int64                  `json:"id"`
	DateOrder  string                 `json:"date_order"`
	Total      float64                `json:"total_value"`
	ClientName string                 `json:"client_name"`
	Products   []OrderProductResponse `json:"products"`
}

type OrderProductResponse struct {
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
}