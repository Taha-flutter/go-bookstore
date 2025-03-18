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
	// fmt.Println(newbook)
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
	bookdetails, _, err := models.GetBookById(ID)
	if err != nil {
		fmt.Println("Error while parsing")

	}
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
	bookId := vars["id"]
	fmt.Println(bookId)

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// fmt.Println("Error while parsing", err)
		utills.ErrorResponse(w, http.StatusInternalServerError, "Error while parsing")
		return
	}
	book, err := models.DeleteBook(ID)

	if err != nil {
		// fmt.Println("Error No book found")
		utills.ErrorResponse(w, http.StatusInternalServerError, "Error No book found")
		return
	}
	res, err := json.Marshal(book)
	if err != nil {
		utills.ErrorResponse(w, http.StatusInternalServerError, "There was an error while preparing the response data")
		return

	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.Book{}
	utills.ParseBody(r, &updateBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookdetails, db, err := models.GetBookById(ID)
	if err != nil {
		fmt.Println("Error while parsing")
	}
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
