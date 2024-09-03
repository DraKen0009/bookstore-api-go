package models

import (
	"github.com/DraKen0009/go-bookstore/pkg/config"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	Publication string `gorm:"type:varchar(100)" json:"publication"`
}

func init() {
	var err error
	err = config.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db = config.GetDB()

	// Automatically migrate the schema
	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate Book model: %v", err)
	}
}

// CreateBook creates a new book record in the database.
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all book records from the database.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById retrieves a book by its ID.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a book by its ID.
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(&book)
	return book
}
