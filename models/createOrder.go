package model

type CreateOrder struct {
	ClientId int64             `json:"clientId"`
	Products []productQuantity `json:"products"`
}

type productQuantity struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}