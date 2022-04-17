package models

type Cart_items_db struct {
	SessionID     string `gorm:"primaryKey"`
	UserID        string `json:"user_id"`
	ProductID     string `json:"productID"`
	Product_name  string
	Product_photo string
	Description   string
	Quantity      string `json:"quantity"`
	CreatedAt     string `json:"created"`
	ModifiedAt    string `json:"modified"`
	StoreID       string `json:"store_id"`
}
