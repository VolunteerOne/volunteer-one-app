package service

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	FindUserFromEmail(string, models.Users) (models.Users, error)
	SaveResetCodeToUser(uuid.UUID, models.Users) error
	ChangePassword(string, models.Users) error
	CreateUser(models.Users) (models.Users, error)
	HashPassword([]byte) ([]byte, error)
}

type loginService struct {
	loginRepository repository.LoginRepository
}

// Instantiated in router.go
func NewLoginService(r repository.LoginRepository) LoginService {
	return loginService{
		loginRepository: r,
	}
}

func (l loginService) FindUserFromEmail(email string, user models.Users) (models.Users, error) {
	return l.loginRepository.FindUserFromEmail(email, user)
}

func (l loginService) SaveResetCodeToUser(resetCode uuid.UUID, user models.Users) error {
	return l.loginRepository.SaveResetCodeToUser(resetCode, user)
}

func (l loginService) ChangePassword(newPassword string, user models.Users) error {
	return l.loginRepository.ChangePassword(newPassword, user)
}

func (l loginService) CreateUser(user models.Users) (models.Users, error) {
	log.Println("[LoginService] Create user...")

	return l.loginRepository.CreateUser(user)
}

func (l loginService) HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return hash, err
}
