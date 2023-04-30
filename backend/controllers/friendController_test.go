package controllers

import (
	"fmt"
	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FriendsControllerUnitTestSuite struct {
	suite.Suite
	friendsObject      models.Friend
	arrayFriendsObject []models.Friend
	c                  *gin.Context
	w                  *httptest.ResponseRecorder
	mockService        *mocks.FriendService
	controller         FriendController
	requestBody        struct {
		FriendOneHandle string
		FriendTwoHandle string
	}
	err     error
	paramID string
}

// Ran before every test
func (suite *FriendsControllerUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockService = new(mocks.FriendService)
	suite.controller = NewFriendController(suite.mockService)

	suite.friendsObject.RelationshipBit = "pending"
	suite.arrayFriendsObject = append(suite.arrayFriendsObject, suite.friendsObject)

	suite.err = fmt.Errorf("error")

	// Used for Reject, Accept, and One
	suite.paramID = "25"
	suite.c.AddParam("id", suite.paramID)
}

// Ran after every test finishes
func (suite *FriendsControllerUnitTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

// Run all the tests in the FriendsControllerUnitTestSuite
func TestFriendsUnitTestSuite(t *testing.T) {
	suite.Run(t, new(FriendsControllerUnitTestSuite))
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Create_BindBodyFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Create_FailCreateFriendService() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateFriend", suite.friendsObject).Return(suite.friendsObject, suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Create_Success() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateFriend", suite.friendsObject).Return(suite.friendsObject, nil)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Reject_OneFriendFail() {
	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, suite.err)
	suite.controller.Reject(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Reject_RejectFriendFail() {
	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, nil)
	suite.mockService.On("RejectFriend", suite.friendsObject).Return(suite.err)
	suite.controller.Reject(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Reject_Success() {
	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, nil)
	suite.mockService.On("RejectFriend", suite.friendsObject).Return(nil)
	suite.controller.Reject(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Accept_OneFriendFail() {
	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, suite.err)
	suite.controller.Accept(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Accept_AcceptFriendFail() {
	suite.friendsObject.RelationshipBit = "friends"

	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, nil)
	suite.mockService.On("AcceptFriend", suite.friendsObject).Return(suite.friendsObject, suite.err)
	suite.controller.Accept(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_Accept_Success() {
	suite.friendsObject.RelationshipBit = "friends"

	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, nil)
	suite.mockService.On("AcceptFriend", suite.friendsObject).Return(suite.friendsObject, nil)
	suite.controller.Accept(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_One_OneFriendFail() {
	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, suite.err)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_One_Success() {
	suite.mockService.On("OneFriend", suite.paramID).Return(suite.friendsObject, nil)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_All_GetFriendsFail() {
	suite.mockService.On("GetFriends").Return(suite.arrayFriendsObject, suite.err)
	suite.controller.All(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *FriendsControllerUnitTestSuite) TestFriendController_All_Success() {
	suite.mockService.On("GetFriends").Return(suite.arrayFriendsObject, nil)
	suite.controller.All(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}
