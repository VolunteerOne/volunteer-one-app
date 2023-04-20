package service

import (
	"testing"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestUsersService_CreateUser(t *testing.T) {
	email := "test@email.com"
	password := "test-password"
	birthdate := "01/01/2000"
	firstname := "test"
	lastname := "user"

	var user models.Users
	user.Email = email
	user.Password = password
	user.Birthdate = birthdate
	user.FirstName = firstname
	user.LastName = lastname

	// new mock repo object
	mockRepo := new(mocks.UsersRepository)
	mockRepo.On("CreateUser", user).Return(user, nil)

	// run actual handler
	fromRepo := NewUsersService(mockRepo)
	res, err := fromRepo.CreateUser(user)

	// checks
	mockRepo.AssertExpectations(t)
	assert.Equal(t, res, user)
	assert.Nil(t, err)
}

func TestUsersService_DeleteUser(t *testing.T) {
	email := "test@email.com"
	password := "test-password"
	birthdate := "01/01/2000"
	firstname := "test"
	lastname := "user"

	var user models.Users
	user.Email = email
	user.Password = password
	user.Birthdate = birthdate
	user.FirstName = firstname
	user.LastName = lastname

	// new mock repo object
	mockRepo := new(mocks.UsersRepository)
	mockRepo.On("CreateUser", user).Return(user, nil)
	mockRepo.On("DeleteUser", user).Return(user, nil)

	// run actual handler
	fromRepo := NewUsersService(mockRepo)
	_, err1 := fromRepo.CreateUser(user)
	res, err := fromRepo.DeleteUser(user)

	_ = err1

	// checks
	mockRepo.AssertExpectations(t)
	assert.Equal(t, res, user)
	assert.Nil(t, err)
}
