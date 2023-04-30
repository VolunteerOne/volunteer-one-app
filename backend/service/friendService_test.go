package service

import (
	"bytes"
	"fmt"
	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type FriendsServiceUnitTestSuite struct {
	suite.Suite
	friendsObject      models.Friend
	arrayFriendsObject []models.Friend
	c                  *gin.Context
	w                  *httptest.ResponseRecorder
	mockRepo           *mocks.FriendRepository
	service            FriendService
	err                error
	paramID            string
}

// Ran before every test
func (suite *FriendsServiceUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockRepo = new(mocks.FriendRepository)
	suite.service = NewFriendService(suite.mockRepo)

	suite.arrayFriendsObject = append(suite.arrayFriendsObject, suite.friendsObject)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
}

// Ran after every test finishes
func (suite *FriendsServiceUnitTestSuite) AfterTest(_, _ string) {
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run all the tests in the FriendsServiceUnitTestSuite
func TestFriendsUnitTestSuite(t *testing.T) {
	suite.Run(t, new(FriendsServiceUnitTestSuite))
}

func (suite *FriendsServiceUnitTestSuite) TestFriendService_CreateFriend() {
	suite.mockRepo.On("CreateFriend", suite.friendsObject).Return(suite.friendsObject, nil)
	res, err := suite.service.CreateFriend(suite.friendsObject)

	assert.Equal(suite.T(), res.FriendOneHandle, "")
	assert.Equal(suite.T(), res.FriendTwoHandle, "")
	assert.Equal(suite.T(), res.RelationshipBit, "")
	assert.Nil(suite.T(), err)
}

func (suite *FriendsServiceUnitTestSuite) TestFriendService_AcceptFriend() {
	suite.mockRepo.On("AcceptFriend", suite.friendsObject).Return(suite.friendsObject, nil)
	res, err := suite.service.AcceptFriend(suite.friendsObject)

	assert.Equal(suite.T(), res.FriendOneHandle, "")
	assert.Equal(suite.T(), res.FriendTwoHandle, "")
	assert.Equal(suite.T(), res.RelationshipBit, "")
	assert.Nil(suite.T(), err)
}

func (suite *FriendsServiceUnitTestSuite) TestFriendService_RejectFriend() {
	suite.mockRepo.On("RejectFriend", suite.friendsObject).Return(nil)
	err := suite.service.RejectFriend(suite.friendsObject)

	assert.Nil(suite.T(), err)
}

func (suite *FriendsServiceUnitTestSuite) TestFriendService_OneFriend() {
	suite.mockRepo.On("OneFriend", suite.paramID).Return(suite.friendsObject, nil)
	res, err := suite.service.OneFriend(suite.paramID)

	assert.Equal(suite.T(), res.FriendOneHandle, "")
	assert.Equal(suite.T(), res.FriendTwoHandle, "")
	assert.Equal(suite.T(), res.RelationshipBit, "")
	assert.Nil(suite.T(), err)
}

func (suite *FriendsServiceUnitTestSuite) TestFriendService_GetFriends() {
	suite.mockRepo.On("GetFriends").Return(suite.arrayFriendsObject, nil)
	res, err := suite.service.GetFriends()

	assert.GreaterOrEqual(suite.T(), len(res), 1)
	assert.Nil(suite.T(), err)
}

func (suite *FriendsServiceUnitTestSuite) TestFriendService_Bind() {
	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	suite.c.Request = req

	var body struct{}
	err := suite.service.Bind(suite.c, &body)

	assert.Nil(suite.T(), err)
}
