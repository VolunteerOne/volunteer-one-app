package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {

	// Get DB Connection values
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	// Create Connection string
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)

	//Connect to the database
	database, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		log.Fatalf("Could not open database.\n") // Handle errors
	}

	log.Printf("Database opened successfully.\n")
	db = database
}

func GetDatabase() *gorm.DB {
	// Return the database object stored
	return db
}
