package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/server"
	"github.com/joho/godotenv"
	// "golang.org/x/text/date"
)

type Users struct {
	gorm.Model
	id       uint   `gorm:"unique; autoincrement;primaryKey"`
	handle   string `gorm: "not null"`
	email    string `gorm: "NOT NULL"`
	password string `gorm: "NOT NULL"`
	// birthdate datatypes.Date `gorm: "NOT NULL"`
	birthdate string `gorm: "NOT NULL"`
	firstName string `gorm: "NOT NULL"`
	lastName  string `gorm: "NOT NULL"`
	// profilePic mediumblob,
	interests string
	verified  uint
}

var db *gorm.DB

func main() {
	insert("yaacii", "ci@gmail.com", "siu", "2020-07-17", "Cicelia", "Siu")
}

func insert(handleA string, emailA string, passowordA string, birthdateA string, firstNameA string, lastNameA string) {
	// database.Init() // Connect database
	URL := fmt.Sprintf("cicelia:siu@tcp(host.docker.internal)/volunteerone?charset=utf8&parseTime=True&loc=Local")

	//Connect to the database
	database, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		log.Fatalf("Could not open database.\n") // Handle errors
	}
	if os.Getenv("DB_MIGRATION") != "" {
		models.Init()
	}

	server.Init() // Start Server

	// var userid string
	// newUser := Users{handle: handleA, email: emailA, password: passowordA, birthdate: birthdateA, firstName: firstNameA, lastName: lastNameA}

	// db := database.GetDatabase()
	// sql := db.Raw("INSERT INTO users (handle, email, password, birthdate, firstName, lastName) VALUES (?, value2, value3, ...);").Scan(&userid)
	// err2 := db.Exec(sql)
	// err2 := db.Select("handle", "email", "password", "birthdate", "firstName", "lastName").Create(&newUser)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	fmt.Println("Done")

	db = database

}

// Runs before main
func init() {
	// Load Environment Variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not load environment variables")
	}

	log.Printf("Application starting on port: %v", os.Getenv("PORT"))
}
