package utils

import (
	"src/src/Carts"
	"src/src/Offers"
	"src/src/Stores"
	"src/src/Users"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("Users.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Users.User3{})
	db.AutoMigrate(Offers.Offer{})
	db.AutoMigrate(&Carts.Cart_items{})
	db.AutoMigrate(&Stores.Store_inventory{})
	DB = db
}
