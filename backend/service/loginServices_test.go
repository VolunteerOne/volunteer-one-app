package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginService_FindUserFromEmail(t *testing.T) {
	email := "test@user.com"

	var user models.User

	var exampleUser models.User
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
