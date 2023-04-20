package controllers

import (
	"bytes"

	"net/http"

	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
)

func TestUserController_CreateSuccess(t *testing.T) {
	// handle := "testHandle"
	// email := "test@email.com"
	// password := "test-password"
	// birthdate := "01/01/2000"
	// firstname := "test"
	// lastname := "user"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// fake := []byte(`{"handle": "testHandle", "email": "test@email.com", "password": "test-password", "firstname": "test", "lastname": "user", "birthdate": "01/01/2000"}`)
	fake := []byte(`{"handle": "what}`)
	req := httptest.NewRequest("POST", "/user/", bytes.NewBuffer(fake))
	c.Request = req

	var user models.Users
	user.Handle = ""
	user.Email = ""
	user.Password = ""
	user.Birthdate = ""
	user.FirstName = ""
	user.LastName = ""

	mockService := new(mocks.UsersService)
	// mockService.On("HashPassword", []byte(user.Password)).Return([]byte(""), nil)
	user.Password = ""

	mockService.On("CreateUser", user).Return(user, nil)

	res := NewUsersController(mockService)

	res.Create(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, c.Writer.Status(), http.StatusOK)
}

func TestUserController_CreateNull(t *testing.T) {
	// email := "test@email.com"
	// password := "test-password"
	// birthdate := "01/01/2000"
	// firstname := "test"
	// lastname := "user"
	// email := ""
	// password := ""
	// birthdate := ""
	// firstname := ""
	// lastname := ""

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	fake := []byte(`{"email": "test@email.com", "password": "test-password", "firstname": "test", "lastname": "user", "birthdate": "01/01/2000"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	c.Request = req

	var user models.Users
	user.Email = ""
	user.Password = ""
	user.Birthdate = ""
	user.FirstName = ""
	user.LastName = ""

	mockService := new(mocks.UsersService)
	// mockService.On("HashPassword", []byte(user.Password)).Return([]byte("hashed pass"), nil)
	user.Password = ""
	// user.Password = """
	mockService.On("CreateUser", user).Return(user, nil)

	res := NewUsersController(mockService)

	res.Create(c)

	mockService.AssertExpectations(t)
	// assert.Equal(t, c.Writer.Status(), http.StatusBadRequest)
	assert.Equal(t, 200, w.Code)
}

func TestUserController_DeleteSuccess(t *testing.T) {
	// handle := "testHandle"
	// email := "test@email.com"
	// password := "test-password"
	// birthdate := "01/01/2000"
	// firstname := "test"
	// lastname := "user"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	fake := []byte(`{"handle": "", "email": "", "password": "", "firstname": "", "": "", "": ""}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	c.Request = req

	var user models.Users
	user.Handle = ""
	user.Email = ""
	user.Password = ""
	user.Birthdate = ""
	user.FirstName = ""
	user.LastName = ""

	mockService := new(mocks.UsersService)
	// mockService.On("HashPassword", []byte(user.Password)).Return([]byte("hashed pass"), nil)
	// user.Password = "hashed pass"
	mockService.On("OneUser", "", user).Return(user, nil)
	mockService.On("DeleteUser", user).Return(user, nil)

	res := NewUsersController(mockService)

	// res.Create(c)
	res.Delete(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, c.Writer.Status(), http.StatusOK)
}
