package models

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
)

type Model interface {
}

// Register all models into this table
var tables = []Model{
	&Object{},
	&User{},
	&Organization{},
	&OrgUsers{},
	&Users{},
}

func Init() {
	// Create migration for all of our tables
	for _, model := range tables {
		log.Printf("Database Migration -> %T", model)
		if database.GetDatabase().AutoMigrate(model) != nil {
			log.Fatalf("Could not complete database migration.\n")
		}
	}
	log.Printf("Database migration successful.\n")
}
