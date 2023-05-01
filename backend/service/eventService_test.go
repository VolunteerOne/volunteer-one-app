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

type EventServiceUnitTestSuite struct {
	suite.Suite
	object      models.Event
	arrayObject []models.Event
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockRepo    *mocks.EventRepository
	service     EventService
	err         error
	paramID     string
}

// Ran before every test
func (suite *EventServiceUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockRepo = new(mocks.EventRepository)
	suite.service = NewEventService(suite.mockRepo)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
}

// Ran after every test finishes
func (suite *EventServiceUnitTestSuite) AfterTest(_, _ string) {
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run all the tests in the EventServiceUnitTestSuite
func TestEventUnitTestSuite(t *testing.T) {
	suite.Run(t, new(EventServiceUnitTestSuite))
}

func (suite *EventServiceUnitTestSuite) TestEventService_CreateEvent() {
	suite.mockRepo.On("CreateEvent", suite.object).Return(suite.object, nil)
	res, err := suite.service.CreateEvent(suite.object)

	assert.Equal(suite.T(), res.OrganizationID, uint(0))
	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Address, "")
	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Description, "")
	assert.Equal(suite.T(), res.Interests, "")
	assert.Equal(suite.T(), res.Skills, "")
	assert.Equal(suite.T(), res.GoodFor, "")
	assert.Equal(suite.T(), res.CauseAreas, "")
	assert.Equal(suite.T(), res.Requirements, "")
	assert.Nil(suite.T(), err)
}

func (suite *EventServiceUnitTestSuite) TestEventService_DeleteEvent() {
	suite.mockRepo.On("DeleteEvent", suite.object).Return(nil)
	suite.err = suite.service.DeleteEvent(suite.object)

	assert.Nil(suite.T(), suite.err)
}

func (suite *EventServiceUnitTestSuite) TestEventService_GetEventById() {
	suite.mockRepo.On("GetEventById", suite.paramID).Return(suite.object, nil)
	res, err := suite.service.GetEventById(suite.paramID)

	assert.Equal(suite.T(), res.OrganizationID, uint(0))
	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Address, "")
	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Description, "")
	assert.Equal(suite.T(), res.Interests, "")
	assert.Equal(suite.T(), res.Skills, "")
	assert.Equal(suite.T(), res.GoodFor, "")
	assert.Equal(suite.T(), res.CauseAreas, "")
	assert.Equal(suite.T(), res.Requirements, "")
	assert.Nil(suite.T(), err)
}

func (suite *EventServiceUnitTestSuite) TestEventService_GetEvents() {
	suite.mockRepo.On("GetEvents").Return(suite.arrayObject, nil)
	res, err := suite.service.GetEvents()

	assert.GreaterOrEqual(suite.T(), len(res), 1)
	assert.Nil(suite.T(), err)
}

func (suite *EventServiceUnitTestSuite) TestEventService_UpdateEvent() {
	suite.mockRepo.On("UpdateEvent", suite.object).Return(suite.object, nil)
	res, err := suite.service.UpdateEvent(suite.object)

	assert.Equal(suite.T(), res.OrganizationID, uint(0))
	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Address, "")
	assert.Equal(suite.T(), res.Name, "")
	assert.Equal(suite.T(), res.Description, "")
	assert.Equal(suite.T(), res.Interests, "")
	assert.Equal(suite.T(), res.Skills, "")
	assert.Equal(suite.T(), res.GoodFor, "")
	assert.Equal(suite.T(), res.CauseAreas, "")
	assert.Equal(suite.T(), res.Requirements, "")
	assert.Nil(suite.T(), err)
}

func (suite *EventServiceUnitTestSuite) TestEventService_Bind() {
	fake := []byte(`{"nope"}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fake))
	suite.c.Request = req

	var body struct{}
	err := suite.service.Bind(suite.c, &body)

	assert.Nil(suite.T(), err)
}
