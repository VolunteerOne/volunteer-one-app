package service

import (
	"bytes"
	"fmt"
	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"net/http/httptest"
	"testing"
)

type UsersServiceUnitTestSuite struct {
	suite.Suite
	object      models.Users
	arrayObject []models.Users
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockRepo    *mocks.UsersRepository
	service     UsersService
	err         error
	paramID     string
	fakeUUID    uuid.UUID
}

// Ran before every test
func (suite *UsersServiceUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockRepo = new(mocks.UsersRepository)
	suite.service = NewUsersService(suite.mockRepo)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
	suite.fakeUUID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
}

// Ran after every test finishes
func (suite *UsersServiceUnitTestSuite) AfterTest(_, _ string) {
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run all the tests in the UsersServiceUnitTestSuite
func TestUsersUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UsersServiceUnitTestSuite))
}

func (suite *UsersServiceUnitTestSuite) TestUsersService_CreateUsers() {
	suite.mockRepo.On("CreateUser", suite.object).Return(suite.object, nil)
	res, err := suite.service.CreateUser(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.Email, "")
	assert.Equal(suite.T(), res.Password, "")
	assert.Equal(suite.T(), res.Birthdate, "")
	assert.Equal(suite.T(), res.FirstName, "")
	assert.Equal(suite.T(), res.LastName, "")
	assert.Equal(suite.T(), res.Interests, "")
	assert.Equal(suite.T(), res.Verified, uint(0))
	assert.Equal(suite.T(), res.ResetCode, suite.fakeUUID)
	assert.Nil(suite.T(), err)
}

func (suite *UsersServiceUnitTestSuite) TestUsersService_OneUser() {
	suite.mockRepo.On("OneUser", suite.paramID, suite.object).Return(suite.object, nil)
	res, err := suite.service.OneUser(suite.paramID, suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.Email, "")
	assert.Equal(suite.T(), res.Password, "")
	assert.Equal(suite.T(), res.Birthdate, "")
	assert.Equal(suite.T(), res.FirstName, "")
	assert.Equal(suite.T(), res.LastName, "")
	assert.Equal(suite.T(), res.Interests, "")
	assert.Equal(suite.T(), res.Verified, uint(0))
	assert.Equal(suite.T(), res.ResetCode, suite.fakeUUID)
	assert.Nil(suite.T(), err)
}

func (suite *UsersServiceUnitTestSuite) TestUsersService_UpdateUsers() {
	suite.mockRepo.On("UpdateUser", suite.object).Return(suite.object, nil)
	res, err := suite.service.UpdateUser(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.Email, "")
	assert.Equal(suite.T(), res.Password, "")
	assert.Equal(suite.T(), res.Birthdate, "")
	assert.Equal(suite.T(), res.FirstName, "")
	assert.Equal(suite.T(), res.LastName, "")
	assert.Equal(suite.T(), res.Interests, "")
	assert.Equal(suite.T(), res.Verified, uint(0))
	assert.Equal(suite.T(), res.ResetCode, suite.fakeUUID)
	assert.Nil(suite.T(), err)
}

func (suite *UsersServiceUnitTestSuite) TestUsersService_DeleteUsers() {
	suite.mockRepo.On("DeleteUser", suite.object).Return(suite.object, nil)
	res, err := suite.service.DeleteUser(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.Email, "")
	assert.Equal(suite.T(), res.Password, "")
	assert.Equal(suite.T(), res.Birthdate, "")
	assert.Equal(suite.T(), res.FirstName, "")
	assert.Equal(suite.T(), res.LastName, "")
	assert.Equal(suite.T(), res.Interests, "")
	assert.Equal(suite.T(), res.Verified, uint(0))
	assert.Equal(suite.T(), res.ResetCode, suite.fakeUUID)
	assert.Nil(suite.T(), err)
}

func (suite *UsersServiceUnitTestSuite) TestUsersService_HashPassword() {
	hash, err := suite.service.HashPassword([]byte("helloworld"))
	err = bcrypt.CompareHashAndPassword(hash, []byte("helloworld"))

	assert.Nil(suite.T(), err)
}

func (suite *UsersServiceUnitTestSuite) TestUsersService_Bind() {
	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	suite.c.Request = req

	var body struct{}
	err := suite.service.Bind(suite.c, &body)

	assert.Nil(suite.T(), err)
}
