package models

import (
	"github.com/jinzhu/gorm"
	"github.com/souravsk/Book_Store_Management_go/pkg/config"
)

var db *gorm.DB

type Book struct { /* Struct are basicle base on models and a models is something that give structer to help as store something into database */
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()    /* calling the connection func to connect the database */
	db = config.GetDB() /* getting all the data */
	db.AutoMigrate(&Book{})
	/* SO db is what that help to talk to database */
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) /* This NewRecord is in gorm it has all the mysql quar so that we don't have to write it that we are using gorm */
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
