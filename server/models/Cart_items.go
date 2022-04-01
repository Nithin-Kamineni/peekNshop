package models

type Cart_items struct {
	UserID     string `json:"id"`
	ProductID  string `json:"productID"`
	Quantity   string `json:"quantity"`
	CreatedAt  string `json:"created"`
	ModifiedAt string `json:"modified"`
}

type UserIDtab struct {
	UserID string `json:"userID"`
}
