package main

import (
	"log"
	"os"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/server"
	"github.com/joho/godotenv"
)

func main() {
	database.Init()
	server.Init()
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
