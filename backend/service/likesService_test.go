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

type LikesServiceUnitTestSuite struct {
	suite.Suite
	object      models.Likes
	arrayObject []models.Likes
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockRepo    *mocks.LikesRepository
	service     LikesService
	err         error
	paramID     string
	likes       int64
}

// Ran before every test
func (suite *LikesServiceUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockRepo = new(mocks.LikesRepository)
	suite.service = NewLikesService(suite.mockRepo)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
}

// Ran after every test finishes
func (suite *LikesServiceUnitTestSuite) AfterTest(_, _ string) {
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run all the tests in the LikesServiceUnitTestSuite
func TestLikesServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, new(LikesServiceUnitTestSuite))
}

func (suite *LikesServiceUnitTestSuite) TestLikesService_CreateLike() {
	suite.mockRepo.On("CreateLike", suite.object).Return(suite.object, nil)
	res, err := suite.service.CreateLike(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostsID, uint(0))
	assert.Nil(suite.T(), err)
}

func (suite *LikesServiceUnitTestSuite) TestLikesService_DeleteLike() {
	suite.mockRepo.On("DeleteLike", suite.object).Return(nil)
	suite.err = suite.service.DeleteLike(suite.object)

	assert.Nil(suite.T(), suite.err)
}

func (suite *LikesServiceUnitTestSuite) TestLikesService_FindLike() {
	suite.mockRepo.On("FindLike", suite.paramID).Return(suite.object, nil)
	res, err := suite.service.FindLike(suite.paramID)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostsID, uint(0))
	assert.Nil(suite.T(), err)
}

func (suite *LikesServiceUnitTestSuite) TestLikesService_AllLikes() {
	suite.mockRepo.On("AllLikes").Return(suite.arrayObject, nil)
	res, err := suite.service.AllLikes()

	assert.GreaterOrEqual(suite.T(), len(res), 1)
	assert.Nil(suite.T(), err)
}

func (suite *LikesServiceUnitTestSuite) TestLikesService_GetLikes() {
	suite.mockRepo.On("GetLikes", suite.paramID).Return(suite.likes, nil)
	res, err := suite.service.GetLikes(suite.paramID)

	assert.Equal(suite.T(), suite.likes, res)
	assert.Nil(suite.T(), err)
}

func (suite *LikesServiceUnitTestSuite) TestLikesService_Bind() {
	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	suite.c.Request = req

	var body struct{}
	err := suite.service.Bind(suite.c, &body)

	assert.Nil(suite.T(), err)
}
