package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/middleware"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/go-gomail/gomail"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginService interface {
	FindUserFromEmail(string, models.Users) (models.Users, error)
	SaveResetCodeToUser(uuid.UUID, models.Users) error
	ChangePassword([]byte, models.Users) error
	HashPassword([]byte) ([]byte, error)
	CompareHashedAndUserPass([]byte, string) error
	GenerateJWT(jwt.SigningMethod, jwt.Claims, string) (string, error)
	GenerateExpiresJWT() (*jwt.NumericDate, *jwt.NumericDate)
	ValidateJWT(string, string) (*jwt.Token, error)
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
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func (l loginService) CompareHashedAndUserPass(hashedPassword []byte, stringPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(stringPassword))
}

func (l loginService) GenerateJWT(
	signingMethod jwt.SigningMethod,
	claims jwt.Claims,
	secret string) (string, error) {

	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString([]byte(secret))
}

func (l loginService) GenerateExpiresJWT() (*jwt.NumericDate, *jwt.NumericDate) {
	// 15 minute expire for accessToken
	accessExpire := jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
	// 1 day expire for refreshToken
	refreshExpire := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	return accessExpire, refreshExpire
}

func (l loginService) ValidateJWT(token string, secret string) (*jwt.Token, error) {
	// hooks into middleware
	return middleware.Validate(token, secret)
}

func (l loginService) MapJWTClaims(token jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
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
