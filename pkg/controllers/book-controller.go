package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Taha-flutter/go-bookstore/pkg/models"
	"github.com/Taha-flutter/go-bookstore/pkg/utills"
	"github.com/gorilla/mux"
)

var newbook []models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newbook := models.GetAllBooks()
	res, _ := json.Marshal(newbook)
	fmt.Println(newbook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")

	}
	bookdetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utills.ParseBody(r, &book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utills.ParseBody(r, &updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookdetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookdetails.Publication = updateBook.Publication
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
