package Carts

type Cart_items struct {
	userID     string `gorm:"primary_key" json:"id"`
	productID  string `json:"productID"`
	quantity   string `json:"quantity"`
	createdAt  string `json:"created"`
	ModifiedAt string `json:"modified"`
}

type UserIDtab struct {
	UserID string `json:"userID"`
}
