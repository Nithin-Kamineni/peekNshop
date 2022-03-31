package models

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
	r  *mux.Router
}
