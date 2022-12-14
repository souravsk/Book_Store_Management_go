package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB /* Created a variable of gorm.db type */
)

func Connect() { /* This func help as to open connection with our database which is mysql */
	d, err := gorm.Open("mysql", "database-1:sovu1612@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d /* here putting into db what ever we recived form in d */
}

func GetDB() *gorm.DB {
	return db /* this func just returing all the db */
}
