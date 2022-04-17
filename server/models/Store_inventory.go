package models

type Store_inventory struct {
	StoreID      string `gorm:"primary_key" json:"id"`
	ProductID    string `json:"productID"`
	ProductName  string `json:"productName"`
	ProductPrice string `json:"price"`
	ProductPhoto string `json:"photo"`
	Description  string `json:"description"`
	Quantity     string `json:"quantity"`
	CreatedAt    string `json:"created"`
	ModifiedAt   string `json:"modified"`
	AccessKey    string `json:"accesskey"`
}

type Stores_Information struct {
	StoreID   string `gorm:"primary_key"`
	AccessKey string
	StoreName string
	Photo_ref string
	Address   string
}

type FavorateStores struct {
	StoreID string
	UserID  string
}
