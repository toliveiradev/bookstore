package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/toliveiradev/bookstore/pkg/model"
	"github.com/toliveiradev/bookstore/pkg/utils"
)

var NewBook model.Book

const contentType = "Content-Type"

const jsonContentType = "application/json"

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := model.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := model.GetBookById(int64(ID))
	res, _ := json.Marshal(book)
	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := model.DeleteBook(int64(ID))
	res, _ := json.Marshal(book)
	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &model.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, db := model.GetBookById(int64(ID))
	if updateBook.Title != "" {
		book.Title = updateBook.Title
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Year != "" {
		book.Year = updateBook.Year
	}
	if updateBook.Publisher != "" {
		book.Publisher = updateBook.Publisher
	}
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
