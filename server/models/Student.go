package models

type Student struct {
	ID       string `gorm:"primary_key" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
