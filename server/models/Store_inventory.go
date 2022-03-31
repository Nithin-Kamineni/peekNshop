package models

type Store_inventory struct {
	StoreID      string `gorm:"primary_key" json:"id"`
	ProductID    string `json:"productID"`
	ProductName  string `json:"productName"`
	ProductPrice string `json:"price"`
	Quantity     string `json:"quantity"`
	CreatedAt    string `json:"created"`
	ModifiedAt   string `json:"modified"`
}
