package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
)

// *****************************************************
// /signup
// Should have: "email", "password", "firstname", "lastname"
// in request body
// *****************************************************
func TestLoginController_SignupSuccessful(t *testing.T) {
	email := "test@email.com"
	password := "test-password"
	firstname := "test"
	lastname := "user"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	fake := []byte(`{"email": "test@email.com", "password": "test-password", "firstname": "test", "lastname": "user"}`)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(fake))
	c.Request = req

	var user models.Users
	user.Email = email
	user.Password = password
	user.FirstName = firstname
	user.LastName = lastname

	mockService := new(mocks.LoginService)
	mockService.On("HashPassword", []byte(user.Password)).Return([]byte("hashed pass"), nil)
	user.Password = "hashed pass"
	mockService.On("CreateUser", user).Return(user, nil)

	res := NewLoginController(mockService)

	res.Signup(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, c.Writer.Status(), http.StatusOK)
}

func TestLoginController_SignupBadRequestBody(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(fake))
	c.Request = req

	mockService := new(mocks.LoginService)

	res := NewLoginController(mockService)

	// the password will be updated
	res.Signup(c)

	mockService.AssertExpectations(t)

	assert.Equal(t, c.Writer.Status(), http.StatusBadRequest)
}

func TestLoginController_SignupHashError(t *testing.T) {
	email := "test@email.com"
	password := "test-password"
	firstname := "test"
	lastname := "user"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	fake := []byte(`{"email": "test@email.com", "password": "test-password", "firstname": "test", "lastname": "user"}`)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(fake))
	c.Request = req

	var user models.Users
	user.Email = email
	user.Password = password
	user.FirstName = firstname
	user.LastName = lastname

	mockService := new(mocks.LoginService)
	mockService.On("HashPassword", []byte(user.Password)).Return([]byte("hashed pass"), fmt.Errorf("Bad Hash"))

	res := NewLoginController(mockService)

	res.Signup(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, c.Writer.Status(), http.StatusBadRequest)
}

func TestLoginController_SignupCreateError(t *testing.T) {
	email := "test@email.com"
	password := "test-password"
	firstname := "test"
	lastname := "user"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	fake := []byte(`{"email": "test@email.com", "password": "test-password", "firstname": "test", "lastname": "user"}`)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(fake))
	c.Request = req

	var user models.Users
	user.Email = email
	user.Password = password
	user.FirstName = firstname
	user.LastName = lastname

	mockService := new(mocks.LoginService)
	mockService.On("HashPassword", []byte(user.Password)).Return([]byte("hashed pass"), nil)
	user.Password = "hashed pass"
	mockService.On("CreateUser", user).Return(user, fmt.Errorf("Create error"))

	res := NewLoginController(mockService)

	res.Signup(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, c.Writer.Status(), http.StatusBadRequest)
}

// *****************************************************
// /login/username/password
// *****************************************************

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
