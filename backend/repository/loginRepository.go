package repository

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginRepository interface {
	FindUserFromEmail(string, models.Users) (models.Users, error)
	SaveResetCodeToUser(uuid.UUID, models.Users) error
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

// Saves a reset code for the user in the DB
func (l loginRepository) SaveResetCodeToUser(resetCode uuid.UUID, user models.Users) error {
	log.Println("[LoginRepository] Save Reset Code...")
	// Update the user's reset code and expiry time
	user.ResetCode = resetCode
	// Save the changes to the database
	return l.DB.Save(&user).Error
}
