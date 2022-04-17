package utils

import (
	"src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("Users.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User3{})
	db.AutoMigrate(&models.Offer{})
	db.AutoMigrate(&models.Orders{})
	db.AutoMigrate(&models.FavorateStoresObj{})
	db.AutoMigrate(&models.Cart_items_db{})
	db.AutoMigrate(&models.Store_inventory{})
	db.AutoMigrate(&models.Stores_Information{})
	DB = db
}

func ConnectTestDatabase() {
	db, err := gorm.Open(sqlite.Open("Users_test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User3{})
	db.AutoMigrate(&models.Offer{})
	db.AutoMigrate(&models.Orders{})
	db.AutoMigrate(&models.FavorateStoresObj{})
	db.AutoMigrate(&models.Cart_items_db{})
	db.AutoMigrate(&models.Store_inventory{})
	db.AutoMigrate(&models.Stores_Information{})
	DB = db
}
