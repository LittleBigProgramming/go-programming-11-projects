package config

import (

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
)

var (
	db *gorm.DB
)

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/bookstore_db?charset=utf8mb4&parseTime=True&loc=Local"
	
	d, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}