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
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	fmt.Println(URL)
	database, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		log.Fatalf("Could not open database.\n")
	}

	log.Printf("Database opened successfully.\n")
	db = database
}

func GetDatabase() *gorm.DB {
	return db
}
