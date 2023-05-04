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
	"time"
)

type EventControllerUnitTestSuite struct {
	suite.Suite
	object      models.Event
	arrayObject []models.Event
	c           *gin.Context
	w           *httptest.ResponseRecorder
	mockService *mocks.EventService
	controller  EventController
	requestBody struct {
		OrganizationID uint
		Name           string
		Address        string
		Date           time.Time
		Description    string
		Interests      string
		Skills         string
		GoodFor        string
		CauseAreas     string
		Requirements   string
	}
	err     error
	paramID string
}

func (suite *EventControllerUnitTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.c, _ = gin.CreateTestContext(suite.w)

	suite.mockService = new(mocks.EventService)
	suite.controller = NewEventController(suite.mockService)

	suite.arrayObject = append(suite.arrayObject, suite.object)

	suite.err = fmt.Errorf("error")

	suite.paramID = "25"
	suite.c.AddParam("id", suite.paramID)
}

func (suite *EventControllerUnitTestSuite) AfterTest(_, _ string) {
	suite.mockService.AssertExpectations(suite.T())
}

func TestEventsControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(EventControllerUnitTestSuite))
}

func (suite *EventControllerUnitTestSuite) TestEventController_All_GetEventsFail() {
	suite.mockService.On("GetEvents").Return(suite.arrayObject, suite.err)
	suite.controller.All(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_All_Success() {
	suite.mockService.On("GetEvents").Return(suite.arrayObject, nil)
	suite.controller.All(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Create_BindBodyFail() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Create_FailCreateEventService() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateEvent", suite.object).Return(suite.object, suite.err)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Create_Success() {
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("CreateEvent", suite.object).Return(suite.object, nil)
	suite.controller.Create(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Delete_GetEventByIdFail() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, suite.err)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Delete_DeleteEventFail() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteEvent", suite.object).Return(suite.err)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Delete_Success() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("DeleteEvent", suite.object).Return(nil)
	suite.controller.Delete(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}

func (suite *EventControllerUnitTestSuite) TestEventController_One_GetEventByIdFail() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, suite.err)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_One_Success() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, nil)
	suite.controller.One(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusAccepted)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Update_GetEventByIdFail() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Update_BindBodyFail() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Update_UpdateEventFail() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("UpdateEvent", suite.object).Return(suite.object, suite.err)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusBadRequest)
}

func (suite *EventControllerUnitTestSuite) TestEventController_Update_Success() {
	suite.mockService.On("GetEventById", suite.paramID).Return(suite.object, nil)
	suite.mockService.On("Bind", suite.c, &suite.requestBody).Return(nil)
	suite.mockService.On("UpdateEvent", suite.object).Return(suite.object, nil)
	suite.controller.Update(suite.c)

	assert.Equal(suite.T(), suite.c.Writer.Status(), http.StatusOK)
}
