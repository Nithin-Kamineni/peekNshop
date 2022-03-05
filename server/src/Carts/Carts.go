package Carts

type Cart struct {
	OrderID  string `json:"orderID"`
	quantity string `json:"quantity"`
	userID   string `json:"userID"`
}
