package main

import (
	"log"
	"os"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/server"
	"github.com/joho/godotenv"
)

func main() {
	database.Init() // Connect database
<<<<<<< HEAD

	if os.Getenv("DB_MIGRATION") != "" {
		models.Init()
	}

	server.Init() // Start Server

=======

	if os.Getenv("DB_MIGRATION") != "" {
		models.Init()
	}

	server.Init() // Start Server
>>>>>>> 4e9ea3a (Update (#5))
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
