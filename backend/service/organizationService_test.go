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

type OrganizationServiceUnitTestSuite struct {
	suite.Suite
	object      models.Organization
	arrayObject []models.Organization
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockRepo    *mocks.OrganizationRepository
	service     OrganizationService
	err         error
	paramID     string
}

func (suite *OrganizationServiceUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockRepo = new(mocks.OrganizationRepository)
	suite.service = NewOrganizationService(suite.mockRepo)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
}

func (suite *OrganizationServiceUnitTestSuite) AfterTest(_, _ string) {
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestOrganizationServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, new(OrganizationServiceUnitTestSuite))
}

func (suite *OrganizationServiceUnitTestSuite) TestOrganizationService_CreateOrganization() {
	suite.mockRepo.On("CreateOrganization", suite.object).Return(suite.object, nil)
	res, err := suite.service.CreateOrganization(suite.object)

	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Description, "")
	assert.Equal(suite.T(), res.Verified, false)
	assert.Equal(suite.T(), res.Interests, "")
	assert.Nil(suite.T(), err)
}

func (suite *OrganizationServiceUnitTestSuite) TestOrganizationService_DeleteOrganization() {
	suite.mockRepo.On("DeleteOrganization", suite.object).Return(nil)
	suite.err = suite.service.DeleteOrganization(suite.object)

	assert.Nil(suite.T(), suite.err)
}

func (suite *OrganizationServiceUnitTestSuite) TestOrganizationService_GetOrganizationById() {
	suite.mockRepo.On("GetOrganizationById", suite.paramID).Return(suite.object, nil)
	res, err := suite.service.GetOrganizationById(suite.paramID)

	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Description, "")
	assert.Equal(suite.T(), res.Verified, false)
	assert.Equal(suite.T(), res.Interests, "")
	assert.Nil(suite.T(), err)
}

func (suite *OrganizationServiceUnitTestSuite) TestOrganizationService_GetOrganizations() {
	suite.mockRepo.On("GetOrganizations").Return(suite.arrayObject, nil)
	res, err := suite.service.GetOrganizations()

	assert.GreaterOrEqual(suite.T(), len(res), 1)
	assert.Nil(suite.T(), err)
}

func (suite *OrganizationServiceUnitTestSuite) TestOrganizationService_UpdateOrganizations() {
	suite.mockRepo.On("UpdateOrganization", suite.object).Return(suite.object, nil)
	res, err := suite.service.UpdateOrganization(suite.object)

	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Description, "")
	assert.Equal(suite.T(), res.Verified, false)
	assert.Equal(suite.T(), res.Interests, "")
	assert.Nil(suite.T(), err)
}

func (suite *OrganizationServiceUnitTestSuite) TestOrganizationService_Bind() {
	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	suite.c.Request = req

	var body struct{}
	err := suite.service.Bind(suite.c, &body)

	assert.Nil(suite.T(), err)
}
