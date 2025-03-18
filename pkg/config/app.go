package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:T@F$0L123456@tcp(localhost:3306)/book_management?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// Open a connection to the MySQL database
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connection established")
}

// GetDB returns the current database connection instance.
func GetDB() *gorm.DB {
	return db
}
