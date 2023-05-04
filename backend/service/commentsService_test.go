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

type CommentsServiceUnitTestSuite struct {
	suite.Suite
	object      models.Comments
	arrayObject []models.Comments
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockRepo    *mocks.CommentsRepository
	service     CommentsService
	err         error
	paramID     string
}

// Ran before every test
func (suite *CommentsServiceUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockRepo = new(mocks.CommentsRepository)
	suite.service = NewCommentsService(suite.mockRepo)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
}

// Ran after every test finishes
func (suite *CommentsServiceUnitTestSuite) AfterTest(_, _ string) {
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run all the tests in the CommentsServiceUnitTestSuite
func TestCommentServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, new(CommentsServiceUnitTestSuite))
}

func (suite *CommentsServiceUnitTestSuite) TestCommentsService_CreateComment() {
	suite.mockRepo.On("CreateComment", suite.object).Return(suite.object, nil)
	res, err := suite.service.CreateComment(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostsID, uint(0))
	assert.Equal(suite.T(), res.CommentDescription, "")
	assert.Nil(suite.T(), err)
}

func (suite *CommentsServiceUnitTestSuite) TestCommentsService_DeleteComment() {
	suite.mockRepo.On("DeleteComment", suite.object).Return(nil)
	suite.err = suite.service.DeleteComment(suite.object)

	assert.Nil(suite.T(), suite.err)
}

func (suite *CommentsServiceUnitTestSuite) TestCommentsService_EditPost() {
	suite.mockRepo.On("EditComment", suite.object).Return(suite.object, nil)
	res, err := suite.service.EditComment(suite.object)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostsID, uint(0))
	assert.Equal(suite.T(), res.CommentDescription, "")
	assert.Nil(suite.T(), err)
}

func (suite *CommentsServiceUnitTestSuite) TestCommentsService_FindComment() {
	suite.mockRepo.On("FindComment", suite.paramID).Return(suite.object, nil)
	res, err := suite.service.FindComment(suite.paramID)

	assert.Equal(suite.T(), res.Handle, "")
	assert.Equal(suite.T(), res.PostsID, uint(0))
	assert.Equal(suite.T(), res.CommentDescription, "")
	assert.Nil(suite.T(), err)
}

func (suite *CommentsServiceUnitTestSuite) TestCommentsService_AllComments() {
	suite.mockRepo.On("AllComments").Return(suite.arrayObject, nil)
	res, err := suite.service.AllComments()

	assert.GreaterOrEqual(suite.T(), len(res), 1)
	assert.Nil(suite.T(), err)
}

func (suite *CommentsServiceUnitTestSuite) TestCommentsService_Bind() {
	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	suite.c.Request = req

	var body struct{}
	err := suite.service.Bind(suite.c, &body)

	assert.Nil(suite.T(), err)
}
