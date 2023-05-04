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

type CommentsControllerUnitTestSuite struct {
	suite.Suite
	object      models.Comments
	arrayObject []models.Comments
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockService *mocks.CommentsService
	controller  CommentsController
	requestBody struct {
		PostsID            uint
		Handle             string
		CommentDescription string
	}
	updateRequestBody struct {
		CommentDescription string
	}
	err     error
	paramID string
}

// Ran before every test
func (suite *CommentsControllerUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockService = new(mocks.CommentsService)
	suite.controller = NewCommentsController(suite.mockService)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
	suite.c.AddParam("id", suite.paramID)
}

// Ran after every test finishes
func (suite *CommentsControllerUnitTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

// Run all the tests in the CommentsControllerUnitTestSuite
func TestCommentsControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(CommentsControllerUnitTestSuite))
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_CreateComment_BindBodyFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.CreateComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_CreateComment_FailCreateComment() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateComment", suite.object).Return(suite.object, suite.err)
	suite.controller.CreateComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_CreateComment_Success() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateComment", suite.object).Return(suite.object, nil)
	suite.controller.CreateComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_DeleteComment_FindCommentFail() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, suite.err)
	suite.controller.DeleteComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_DeleteComment_DeleteCommentFail() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteComment", suite.object).Return(suite.err)
	suite.controller.DeleteComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_DeleteComment_Success() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteComment", suite.object).Return(nil)
	suite.controller.DeleteComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_EditComment_FindPostFail() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, suite.err)
	suite.controller.EditComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_EditComment_BindBodyFail() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.updateRequestBody).Return(suite.err)
	suite.controller.EditComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_EditComment_EditCommentFail() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.updateRequestBody).Return(nil)
	suite.mockService.On("EditComment", suite.object).Return(suite.object, suite.err)
	suite.controller.EditComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_EditComment_Success() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.updateRequestBody).Return(nil)
	suite.mockService.On("EditComment", suite.object).Return(suite.object, nil)
	suite.controller.EditComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_FindComment_FindCommentFail() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, suite.err)
	suite.controller.FindComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_FindComment_Success() {
	suite.mockService.On("FindComment", suite.paramID).Return(suite.object, nil)
	suite.controller.FindComment(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_AllComments_AllCommentsFail() {
	suite.mockService.On("AllComments").Return(suite.arrayObject, suite.err)
	suite.controller.AllComments(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *CommentsControllerUnitTestSuite) TestCommentsController_AllComments_Success() {
	suite.mockService.On("AllComments").Return(suite.arrayObject, nil)
	suite.controller.AllComments(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}
