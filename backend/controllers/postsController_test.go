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

type PostsControllerUnitTestSuite struct {
	suite.Suite
	object      models.Posts
	arrayObject []models.Posts
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockService *mocks.PostsService
	controller  PostsController
	requestBody struct {
		Handle          string
		PostDescription string
	}
	updateRequestBody struct {
		PostDescription string
	}
	err     error
	paramID string
}

// Ran before every test
func (suite *PostsControllerUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockService = new(mocks.PostsService)
	suite.controller = NewPostsController(suite.mockService)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	// Used for Reject, Accept, and One
	suite.paramID = "25"
	suite.c.AddParam("id", suite.paramID)
}

// Ran after every test finishes
func (suite *PostsControllerUnitTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

// Run all the tests in the PostsControllerUnitTestSuite
func TestPostsControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(PostsControllerUnitTestSuite))
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_CreatePost_BindBodyFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.CreatePost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_CreatePost_FailCreatePostService() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreatePost", suite.object).Return(suite.object, suite.err)
	suite.controller.CreatePost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_CreatePost_Success() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreatePost", suite.object).Return(suite.object, nil)
	suite.controller.CreatePost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_DeletePost_FindPostFail() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, suite.err)
	suite.controller.DeletePost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_DeletePost_DeletePostFail() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeletePost", suite.object).Return(suite.err)
	suite.controller.DeletePost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostController_DeletePost_Success() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeletePost", suite.object).Return(nil)
	suite.controller.DeletePost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_EditPost_FindPostFail() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, suite.err)
	suite.controller.EditPost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_EditPost_BindBodyFail() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.updateRequestBody).Return(suite.err)
	suite.controller.EditPost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_EditPost_EditPostFail() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.updateRequestBody).Return(nil)
	suite.mockService.On("EditPost", suite.object).Return(suite.object, suite.err)
	suite.controller.EditPost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_EditPost_Success() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.updateRequestBody).Return(nil)
	suite.mockService.On("EditPost", suite.object).Return(suite.object, nil)
	suite.controller.EditPost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_FindPost_FindPostFail() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, suite.err)
	suite.controller.FindPost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_FindPost_Success() {
	suite.mockService.On("FindPost", suite.paramID).Return(suite.object, nil)
	suite.controller.FindPost(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_AllPosts_AllPostsFail() {
	suite.mockService.On("AllPosts").Return(suite.arrayObject, suite.err)
	suite.controller.AllPosts(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *PostsControllerUnitTestSuite) TestPostsController_AllPosts_Success() {
	suite.mockService.On("AllPosts").Return(suite.arrayObject, nil)
	suite.controller.AllPosts(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}
