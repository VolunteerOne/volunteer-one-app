package service

import (
	"testing"

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
	fromRepo := NewLoginService(mockRepo)
	res, err := fromRepo.FindUserFromEmail(email, user)

	// checks
	mockRepo.AssertExpectations(t)
	assert.Equal(t, res, exampleUser)
	assert.Nil(t, err)
}

func TestLoginService_CreateUser(t *testing.T) {
	email := "test@email.com"
	password := "test-password"
	firstname := "test"
	lastname := "user"

	var user models.Users
	user.Email = email
	user.Password = password
	user.FirstName = firstname
	user.LastName = lastname

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	mockRepo.On("CreateUser", user).Return(user, nil)

	// run actual handler
	fromRepo := NewLoginService(mockRepo)
	res, err := fromRepo.CreateUser(user)

	// checks
	mockRepo.AssertExpectations(t)
	assert.Equal(t, res, user)
	assert.Nil(t, err)
}

func TestLoginService_HashPassword(t *testing.T) {
	password := "mypass"

	// new mock repo object
	mockRepo := new(mocks.LoginRepository)
	fromRepo := NewLoginService(mockRepo)
	res, err := fromRepo.HashPassword([]byte(password))

	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}
