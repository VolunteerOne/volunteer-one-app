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
	ChangePassword([]byte, models.Users) error
	FindTokenFromID(uint, models.Delegations) (models.Delegations, error)
	SaveRefreshToken(uint, string, models.Delegations) error
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

func (l loginRepository) ChangePassword(newPassword []byte, user models.Users) error {
	log.Println("Entering ChangePassword repository")

	user.Password = string(newPassword)
	return l.DB.Save(&user).Error
}

func (l loginRepository) FindTokenFromID(userid uint, deleg models.Delegations) (models.Delegations, error) {
	log.Println("[LoginRepository] Find Delegation...")

	err := l.DB.Where("users_id = ?", userid).First(&deleg).Error
	return deleg, err
}

func (l loginRepository) SaveRefreshToken(userid uint, refreshToken string, deleg models.Delegations) error {
	log.Println("[LoginRepository] Save Refresh Token...")

	deleg.UsersID = userid
	deleg.RefreshToken = refreshToken

	err := l.DB.Save(&deleg).Error

	if err != nil {
        log.Println("[LoginRepository:SaveRefreshToken] Have To Update Token Only...")
        return l.DB.Model(&deleg).Where("users_id = ?", userid).Update("refresh_token", refreshToken).Error
	}

	return err
}
