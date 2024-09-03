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

func Connect() error {
	// Define the DSN (Data Source Name) for the MySQL connection.
	dsn := "root:mysql@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the database using the GORM v2 API.
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to the database!")
		return err
	}

	// Assign the connection to the global variable `db`.
	db = d
	fmt.Println("Successfully connected to the database.")
	return nil
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
