package models

type Cart_items_db struct {
	UserID     string `gorm:"-" json:"id"`
	ProductID  string `json:"productID"`
	Quantity   string `json:"quantity"`
	CreatedAt  string `json:"created"`
	ModifiedAt string `json:"modified"`
}
