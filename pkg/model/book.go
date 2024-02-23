package model

import (
	"github.com/jinzhu/gorm"
	"github.com/toliveiradev/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title     string `gorm:""json:"title"`
	Author    string `json:"author"`
	Year      string `json:"year"`
	Publisher string `json:"publisher"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
