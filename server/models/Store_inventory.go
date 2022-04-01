package models

type Store_inventory struct {
	StoreID      string `gorm:"primary_key" json:"id"`
	ProductID    string `json:"productID"`
	ProductName  string `json:"productName"`
	ProductPrice string `json:"price"`
	ProductDesc  string `json:"description"`
	Quantity     string `json:"quantity"`
	CreatedAt    string `json:"created"`
	ModifiedAt   string `json:"modified"`
	AccessKey    string `json:"accesskey"`
}

type Stores_Information struct {
	storeID   string
	StoreName string
	Photo_ref string
	Address   string
}

type FavorateStores struct {
	StoreID string
	UserID  string
}
