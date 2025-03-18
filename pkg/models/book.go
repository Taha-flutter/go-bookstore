package models

import (
	"github.com/Taha-flutter/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books

}
func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	var book Book
	err := db.Where("ID=?", Id).Find(&book).Error
	if err != nil {
		return &book, db, err
	}
	return &book, db, nil

}

func DeleteBook(Id int64) (Book, error) {
	var book Book
	err := db.Where("ID = ?", Id).First(&book).Error
	if err != nil {
		return book, err
	}
	err = db.Delete(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}
