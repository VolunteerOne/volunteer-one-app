package repository

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type LoginRepository interface {
	FindUserFromEmail(string, models.Users) (models.Users, error)
    CreateUser(models.Users) (models.Users, error)
}

type loginRepository struct {
	DB *gorm.DB
}

// Instantiated in router.go
func NewLoginRepository(db *gorm.DB) LoginRepository {
	return loginRepository{
		DB: db,
	}
}

// Attempts to find the first entry with the email in the DB
func (l loginRepository) FindUserFromEmail(email string, user models.Users) (models.Users, error) {
	log.Println("[LoginRepository] Find Email...")

	// User will be populated with the content if possible
	err := l.DB.Where("email = ?", email).First(&user).Error

	return user, err
}

// Add the user to the DB
func (l loginRepository) CreateUser(user models.Users) (models.Users, error) {
    log.Println("[LoginRepository] Create user...")

    // User will be created
    err := l.DB.Create(&user).Error 
 
    return user, err
}
