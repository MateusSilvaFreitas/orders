package model

type Order struct {
	ID        int64   `json:"id"`
	DateOrder string  `json:"date_order"`
	Total     float64 `json:"total_value"`
	ClientID  int64   `json:"client_id"`
}