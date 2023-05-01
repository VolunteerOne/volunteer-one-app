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

type LikesControllerUnitTestSuite struct {
	suite.Suite
	object      models.Likes
	arrayObject []models.Likes
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockService *mocks.LikesService
	controller  LikesController
	requestBody struct {
		Handle  string
		PostsID uint
	}
	err     error
	paramID string
	likes   int64
}

// Ran before every test
func (suite *LikesControllerUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockService = new(mocks.LikesService)
	suite.controller = NewLikesController(suite.mockService)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
	suite.c.AddParam("id", suite.paramID)
}

// Ran after every test finishes
func (suite *LikesControllerUnitTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

// Run all the tests in the LikesControllerUnitTestSuite
func TestLikesControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(LikesControllerUnitTestSuite))
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_CreateLike_BindBodyFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.CreateLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_CreateLike_FailCreateLike() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateLike", suite.object).Return(suite.object, suite.err)
	suite.controller.CreateLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_CreateLike_Success() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateLike", suite.object).Return(suite.object, nil)
	suite.controller.CreateLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_DeleteLike_FindLikeFail() {
	suite.mockService.On("FindLike", suite.paramID).Return(suite.object, suite.err)
	suite.controller.DeleteLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_DeleteLike_DeleteLikeFail() {
	suite.mockService.On("FindLike", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteLike", suite.object).Return(suite.err)
	suite.controller.DeleteLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_DeleteLike_Success() {
	suite.mockService.On("FindLike", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteLike", suite.object).Return(nil)
	suite.controller.DeleteLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_FindLikes_FindLikeFail() {
	suite.mockService.On("FindLike", suite.paramID).Return(suite.object, suite.err)
	suite.controller.FindLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_FindLikes_Success() {
	suite.mockService.On("FindLike", suite.paramID).Return(suite.object, nil)
	suite.controller.FindLike(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_AllLikes_AllLikesFail() {
	suite.mockService.On("AllLikes").Return(suite.arrayObject, suite.err)
	suite.controller.AllLikes(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_AllLikes_Success() {
	suite.mockService.On("AllLikes").Return(suite.arrayObject, nil)
	suite.controller.AllLikes(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_GetLikes_GetLikesFail() {
	suite.mockService.On("GetLikes", suite.paramID).Return(suite.likes, suite.err)
	suite.controller.GetLikes(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *LikesControllerUnitTestSuite) TestLikesController_GetLikes_Success() {
	suite.mockService.On("GetLikes", suite.paramID).Return(suite.likes, nil)
	suite.controller.GetLikes(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}
