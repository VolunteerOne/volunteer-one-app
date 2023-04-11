package service

import (
	"time"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	FindUserFromEmail(string, models.Users) (models.Users, error)
	SaveResetCodeToUser(uuid.UUID, models.Users) error
	ChangePassword([]byte, models.Users) error
	HashPassword([]byte) ([]byte, error)
	CompareHashedAndUserPass([]byte, string) error
	GenerateJWT(uint, time.Time, string) (string, error)
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

func (l loginService) ChangePassword(newPassword []byte, user models.Users) error {
	return l.loginRepository.ChangePassword(newPassword, user)
}

func (l loginService) HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return hash, err
}

func (l loginService) CompareHashedAndUserPass(hashedPassword []byte, stringPassword string) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(stringPassword))
	return err
}

func (l loginService) GenerateJWT(userid uint, exp time.Time, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userid,
		"exp": exp,
	})
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
