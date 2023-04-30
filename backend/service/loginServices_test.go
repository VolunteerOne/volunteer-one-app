package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"testing"
	"time"

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

	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)
	res, err := loginService.HashPassword([]byte(password))

	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func TestLoginService_CompareHashedAndUserPass(t *testing.T) {
	password := "mypass"

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
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	claims := jwt.MapClaims{}

	_, err := loginService.GenerateJWT(jwt.SigningMethodRS384, claims, "")

	assert.NotNil(t, err)
}

func TestLoginService_GoodJWTSigning(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	claims := jwt.MapClaims{}

	token, err := loginService.GenerateJWT(jwt.SigningMethodHS256, claims, "")

	assert.EqualValues(t, token, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.LwimMJA3puF3ioGeS-tfczR3370GXBZMIL-bdpu4hOU")
	assert.Nil(t, err)
}

func TestLoginService_GenerateExpiresJWT(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	accessExpire, refreshExpire := loginService.GenerateExpiresJWT()

	diffAccess := accessExpire.Sub(time.Now())
	assert.LessOrEqual(t, diffAccess.Minutes(), float64(15))

	diffRefresh := refreshExpire.Sub(time.Now())
	assert.LessOrEqual(t, diffRefresh.Hours(), float64(24))
}

func TestLoginService_ValidateJWT(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "admin",
	})
	tokenString, _ := token.SignedString([]byte(""))

	returnedToken, error := loginService.ValidateJWT(tokenString, "")

	assert.True(t, returnedToken.Valid)
	assert.Nil(t, error)
}

func TestLoginService_ValidateJWTError(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "admin",
	})
	tokenString, _ := token.SignedString([]byte(""))

	_, error := loginService.ValidateJWT(tokenString, "")

	assert.NotNil(t, error)
}

func TestLoginService_MapJWTClaims(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)
	loginService := NewLoginService(mockRepo)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "admin",
	})

	res, ok := loginService.MapJWTClaims(*token)

	assert.IsType(t, jwt.MapClaims{}, res)
	assert.True(t, ok)
}

func TestLoginService_SaveRefreshToken(t *testing.T) {
	var d models.Delegations

	mockRepo := new(mocks.LoginRepository)
	mockRepo.On("SaveRefreshToken", uint(0), "", d).Return(nil)

	// run actual handler
	loginService := NewLoginService(mockRepo)
	err := loginService.SaveRefreshToken(uint(0), "", d)

	// checks
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestLoginService_FindRefreshToken(t *testing.T) {
	var d models.Delegations

	mockRepo := new(mocks.LoginRepository)
	mockRepo.On("FindRefreshToken", float64(0), d).Return(d, nil)

	loginService := NewLoginService(mockRepo)
	res, err := loginService.FindRefreshToken(float64(0), d)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, res, d)
	assert.Nil(t, err)
}

func TestLoginService_DeleteRefreshToken(t *testing.T) {
	var d models.Delegations

	mockRepo := new(mocks.LoginRepository)
	mockRepo.On("DeleteRefreshToken", d).Return(nil)

	loginService := NewLoginService(mockRepo)
	err := loginService.DeleteRefreshToken(d)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestLoginService_ParseUUID(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)

	loginService := NewLoginService(mockRepo)
	ans, err := loginService.ParseUUID("00000000-0000-0000-0000-000000000000")

	assert.IsType(t, uuid.UUID{}, ans)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestLoginService_ParseUUIDFail(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)

	loginService := NewLoginService(mockRepo)
	_, err := loginService.ParseUUID("00-0000-0000-0000-000000000000")

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, err)
}

func TestLoginService_GenerateUUID(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)

	loginService := NewLoginService(mockRepo)

	res := loginService.GenerateUUID()

	assert.IsType(t, uuid.UUID{}, res)
}

func TestLoginService_SendResetCodeToEmail(t *testing.T) {
	mockRepo := new(mocks.LoginRepository)

	loginService := NewLoginService(mockRepo)

	err := loginService.SendResetCodeToEmail("", "")

	assert.NotNil(t, err)
}
