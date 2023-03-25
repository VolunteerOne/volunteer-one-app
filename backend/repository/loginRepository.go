package repository

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type LoginRepository interface {
	FindUserFromEmail(string, models.User) (models.User, error)
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
func (l loginRepository) FindUserFromEmail(email string, user models.User) (models.User, error) {
	log.Println("[LoginRepository] Find Email...")

	// User will be populated with the content if possible
	err := l.DB.Where("email = ?", email).First(&user).Error

	return user, err
}
