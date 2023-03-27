package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
)

// Tests when a good email and password occurs
func TestEmailFound(t *testing.T) {
	email := "test@user.com"
	password := "password"

	w := httptest.NewRecorder()
	w.Code = 200 // expected code
	// start new gin context to pass in
	c, _ := gin.CreateTestContext(w)

	c.AddParam("email", email)
	c.AddParam("password", password)

	// example user model to pass in empty
	var emptyUser models.Users

	// expected user model
	var user models.Users
	user.Email = email
	user.Password = password

	// setup mock
	mockService := new(mocks.LoginService)
	// mock the function
	mockService.On("FindUserFromEmail", email, emptyUser).Return(user, nil)

	// run actual handler
	res := NewLoginController(mockService)
	res.Login(c)

	// check that everything happened as expected
	mockService.AssertExpectations(t)

	// Verify response code
	assert.Equal(t, 200, w.Code)
}

// Tests that an 502 is returned when an error is returned from the database
func TestRetrievalError(t *testing.T) {
	email := "test@user.com"
	password := "password"

	w := httptest.NewRecorder()
	w.Code = 400 // expected code
	c, _ := gin.CreateTestContext(w)

	c.AddParam("email", email)
	c.AddParam("password", password)

	var emptyUser models.Users

	var user models.Users
	user.Email = email
	user.Password = password

	mockService := new(mocks.LoginService)
	mockService.On("FindUserFromEmail", email, emptyUser).Return(user, fmt.Errorf("Arrrrr"))

	res := NewLoginController(mockService)
	res.Login(c)

	mockService.AssertExpectations(t)

	assert.Equal(t, 502, w.Code)
}

// Tests that the passed param password and db passwords are different
func TestPasswordsDontMatch(t *testing.T) {
	email := "test@user.com"
	password := "password"

	w := httptest.NewRecorder()
	w.Code = 400
	c, _ := gin.CreateTestContext(w)

	c.AddParam("email", email)
	c.AddParam("password", "not right password")

	var emptyUser models.Users

	var user models.Users
	user.Email = email
	user.Password = password

	mockService := new(mocks.LoginService)
	mockService.On("FindUserFromEmail", email, emptyUser).Return(user, nil)

	res := NewLoginController(mockService)
	res.Login(c)

	mockService.AssertExpectations(t)

	assert.Equal(t, 400, w.Code)
}

