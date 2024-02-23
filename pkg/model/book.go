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

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", ID).Find(&book)
	return &book, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
