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

type PostsServiceUnitTestSuite struct {
	suite.Suite
	object      models.Posts
	arrayObject []models.Posts
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockRepo    *mocks.PostsRepository
	service     PostsService
	err         error
	paramID     string
}

// Ran before every test
func (suite *PostsServiceUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockRepo = new(mocks.PostsRepository)
	suite.service = NewPostsService(suite.mockRepo)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
}

// Ran after every test finishes
func (suite *PostsServiceUnitTestSuite) AfterTest(_, _ string) {
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run all the tests in the PostsServiceUnitTestSuite
func TestPostsServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, new(PostsServiceUnitTestSuite))
}

func (suite *PostsServiceUnitTestSuite) TestPostsService_CreatePost() {
	suite.mockRepo.On("CreatePost", suite.object).Return(suite.object, nil)
	res, err := suite.service.CreatePost(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostDescription, "")
	assert.Equal(suite.T(), res.Likes, uint(0))
	assert.Nil(suite.T(), err)
}

func (suite *PostsServiceUnitTestSuite) TestPostsService_DeletePost() {
	suite.mockRepo.On("DeletePost", suite.object).Return(nil)
	suite.err = suite.service.DeletePost(suite.object)

	assert.Nil(suite.T(), suite.err)
}

func (suite *PostsServiceUnitTestSuite) TestPostsService_EditPost() {
	suite.mockRepo.On("EditPost", suite.object).Return(suite.object, nil)
	res, err := suite.service.EditPost(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostDescription, "")
	assert.Equal(suite.T(), res.Likes, uint(0))
	assert.Nil(suite.T(), err)
}

func (suite *PostsServiceUnitTestSuite) TestPostsService_FindPost() {
	suite.mockRepo.On("FindPost", suite.paramID).Return(suite.object, nil)
	res, err := suite.service.FindPost(suite.paramID)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostDescription, "")
	assert.Equal(suite.T(), res.Likes, uint(0))
	assert.Nil(suite.T(), err)
}

func (suite *PostsServiceUnitTestSuite) TestPostsService_AllPosts() {
	suite.mockRepo.On("AllPosts").Return(suite.arrayObject, nil)
	res, err := suite.service.AllPosts()

	assert.GreaterOrEqual(suite.T(), len(res), 1)
	assert.Nil(suite.T(), err)
}

func (suite *PostsServiceUnitTestSuite) TestPostsService_Bind() {
	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	suite.c.Request = req

	var body struct{}
	err := suite.service.Bind(suite.c, &body)

	assert.Nil(suite.T(), err)
}
