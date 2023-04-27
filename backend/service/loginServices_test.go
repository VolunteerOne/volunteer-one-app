package service

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestLoginService_FindUserFromEmail(t *testing.T) {
	email := "test@user.com"

	var user models.Users

	var exampleUser models.Users
	exampleUser.Email = email
	exampleUser.Password = "password"

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	mockRepo.On("FindUserFromEmail", email, user).Return(exampleUser, nil)

	// run actual handler
	loginService := NewLoginService(mockRepo)
	res, err := loginService.FindUserFromEmail(email, user)

	// checks
	mockRepo.AssertExpectations(t)
	assert.Equal(t, res, exampleUser)
	assert.Nil(t, err)
}

func TestLoginService_SaveResetCodeToUser(t *testing.T) {
	var user models.Users
	fakeCode := uuid.New()

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	mockRepo.On("SaveResetCodeToUser", fakeCode, user).Return(nil)

	// run actual handler
	loginService := NewLoginService(mockRepo)
	err := loginService.SaveResetCodeToUser(fakeCode, user)

	// checks
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestLoginService_ChangePassword(t *testing.T) {
	var user models.Users
	fakePassword := []byte("")

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	mockRepo.On("ChangePassword", fakePassword, user).Return(nil)

	// run actual handler
	loginService := NewLoginService(mockRepo)
	err := loginService.ChangePassword(fakePassword, user)

	// checks
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}
func TestLoginService_HashPassword(t *testing.T) {
	password := "mypass"

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)
	res, err := loginService.HashPassword([]byte(password))

	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func TestLoginService_CompareHashedAndUserPass(t *testing.T) {
	password := "mypass"

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	// generate a hash
	res, err := loginService.HashPassword([]byte(password))

	if err != nil {
		t.Errorf("Hash Password didn't work")
	}

	err = loginService.CompareHashedAndUserPass(res, password)

	assert.Nil(t, err)
}

func TestLoginService_ErrorWhenSigningToken(t *testing.T) {
	// accessExpire := jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
	// refreshExpire := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	claims := jwt.MapClaims{}

	_, err := loginService.GenerateJWT(jwt.SigningMethodRS384, claims, "")

	assert.NotNil(t, err)
}
