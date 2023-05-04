package controllers

import (
	"fmt"
	"github.com/stretchr/testify/suite"

	"net/http"

	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
)

type UsersControllerUnitTestSuite struct {
	suite.Suite
	object      models.Users
	arrayObject []models.Users
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockService *mocks.UsersService
	controller  UsersController
	requestBody struct {
		Id        uint
		Handle    string
		Email     string
		Password  string
		Birthdate string
		FirstName string
		LastName  string
		Interests string
		Verified  uint
	}
	err     error
	paramID string
}

func (suite *UsersControllerUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockService = new(mocks.UsersService)
	suite.controller = NewUsersController(suite.mockService)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	// Used for Reject, Accept, and One
	suite.paramID = "25"
	suite.c.AddParam("id", suite.paramID)
}

// Ran after every test finishes
func (suite *UsersControllerUnitTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

// Run all the tests in the UsersControllerUnitTestSuite
func TestUsersControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UsersControllerUnitTestSuite))
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Create_BindBodyFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Create_HashPasswordFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("HashPassword", []byte("")).Return([]byte(""), suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Create_FailCreateUserService() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("HashPassword", []byte("")).Return([]byte(""), nil)
	suite.mockService.On("CreateUser", suite.object).Return(suite.object, suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Create_Success() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("HashPassword", []byte("")).Return([]byte(""), nil)
	suite.mockService.On("CreateUser", suite.object).Return(suite.object, nil)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_One_OneUsersFail() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, suite.err)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_One_Success() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, nil)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Update_OneUsersFail() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Update_BindBodyFail() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Update_UpdateUserFail() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("UpdateUser", suite.object).Return(suite.object, suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Update_Success() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("UpdateUser", suite.object).Return(suite.object, nil)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Delete_OneUserFail() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, suite.err)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Delete_DeleteUserFail() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, nil)
	suite.mockService.On("DeleteUser", suite.object).Return(suite.object, suite.err)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *UsersControllerUnitTestSuite) TestUsersController_Delete_Success() {
	suite.mockService.On("OneUser", suite.paramID, suite.object).Return(suite.object, nil)
	suite.mockService.On("DeleteUser", suite.object).Return(suite.object, nil)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}
