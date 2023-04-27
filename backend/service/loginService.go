package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginService interface {
	FindUserFromEmail(string, models.Users) (models.Users, error)
	SaveResetCodeToUser(uuid.UUID, models.Users) error
	ChangePassword([]byte, models.Users) error
	HashPassword([]byte) ([]byte, error)
	CompareHashedAndUserPass([]byte, string) error
	GenerateJWT(uint, *jwt.NumericDate, *jwt.NumericDate, string, *gin.Context) (string, string, error)
	SaveRefreshToken(uint, string, models.Delegations) error
	FindRefreshToken(float64, models.Delegations) (models.Delegations, error)
	DeleteRefreshToken(models.Delegations) error
	ParseUUID(string) (uuid.UUID, error)
	MapJWTClaims(jwt.Token) (jwt.MapClaims, bool)
	GenerateUUID() uuid.UUID
	SendResetCodeToEmail(string, string) error
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

func (l loginService) GenerateJWT(userid uint,
	accessExp *jwt.NumericDate,
	refreshExp *jwt.NumericDate,
	secret string,
	c *gin.Context) (string, string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userid,
		"exp":  accessExp,
		"type": "access",
	})
	accessTokenString, err := accessToken.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create access token",
			"success": false,
		})
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userid,
		"exp":  refreshExp,
		"type": "refresh",
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create refresh token",
			"success": false,
		})
		return "", "", err
	}

	return accessTokenString, refreshTokenString, err
}

func (l loginService) SaveRefreshToken(userid uint, refreshToken string, deleg models.Delegations) error {
	return l.loginRepository.SaveRefreshToken(userid, refreshToken, deleg)
}

func (l loginService) FindRefreshToken(userid float64, deleg models.Delegations) (models.Delegations, error) {
	return l.loginRepository.FindRefreshToken(userid, deleg)
}

func (l loginService) DeleteRefreshToken(deleg models.Delegations) error {
	return l.loginRepository.DeleteRefreshToken(deleg)
}

func (l loginService) ParseUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func (l loginService) MapJWTClaims(token jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

func (l loginService) GenerateUUID() uuid.UUID {
	return uuid.New()
}

func (l loginService) SendResetCodeToEmail(email string, resetCode string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "edwardsung4217@gmail.com") //need to replace this with proper volunteer email
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Password Reset Code")
	mailer.SetBody("text/plain", "Your password reset code is "+resetCode)
	return gomail.NewDialer("smtp.sendgrid.net", 465, "apikey", "APIKEY").DialAndSend(mailer)
}
