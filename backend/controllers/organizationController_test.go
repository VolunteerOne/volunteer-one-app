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

type OrganizationControllerUnitTestSuite struct {
	suite.Suite
	object      models.Organization
	arrayObject []models.Organization
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockService *mocks.OrganizationService
	controller  OrganizationController
	requestBody struct {
		Name        string
		Description string
		Verified    bool
		Interests   string
	}
	err     error
	paramID string
}

func (suite *OrganizationControllerUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockService = new(mocks.OrganizationService)
	suite.controller = NewOrganizationController(suite.mockService)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
	suite.c.AddParam("id", suite.paramID)
}

func (suite *OrganizationControllerUnitTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

func TestOrganizationControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(OrganizationControllerUnitTestSuite))
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Create_BindBodyFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Create_FailCreateOrganizationService() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateOrganization", suite.object).Return(suite.object, suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Create_Success() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateOrganization", suite.object).Return(suite.object, nil)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_All_GetOrganizationsFail() {
	suite.mockService.On("GetOrganizations").Return(suite.arrayObject, suite.err)
	suite.controller.All(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_All_Success() {
	suite.mockService.On("GetOrganizations").Return(suite.arrayObject, nil)
	suite.controller.All(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_One_GetOrganizationByIdFail() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, suite.err)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_One_Success() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, nil)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusAccepted)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Update_GetOrganizationByIdFail() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Update_BindBodyFail() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Update_UpdateOrganizationFail() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("UpdateOrganization", suite.object).Return(suite.object, suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Update_Success() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("UpdateOrganization", suite.object).Return(suite.object, nil)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Delete_GetOrganizationByIdFail() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, suite.err)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Delete_DeleteOrganizationFail() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteOrganization", suite.object).Return(suite.err)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *OrganizationControllerUnitTestSuite) TestOrganizationController_Delete_Success() {
	suite.mockService.On("GetOrganizationById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteOrganization", suite.object).Return(nil)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}
