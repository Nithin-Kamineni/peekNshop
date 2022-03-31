package models

type Offer struct {
	Name        string `gorm:"primary_key" json:"name"`
	Description string `json:"description"`
}
