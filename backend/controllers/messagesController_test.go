package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// *****************************************************
// Tests everything in MessagesController >:)
// *****************************************************

// CreateMessage tests

// This test ensures that a proper service call
// returns a success for a message that was created.
func TestMessagesController_CreateMessageSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mockMessageStruct struct {
		FromUserId uint
		ToUserId   uint
		Subject    string
		Message    string
	}

	// The mock message we send & expect as JSON
	mockMessage := mockMessageStruct{
		FromUserId: 1,
		ToUserId:   2,
		Subject:    "...hello, there.",
		Message:    "...gEneRAl kEnoBi...",
	}

	mockResponse := models.Messages{
		FromUsersID: mockMessage.FromUserId,
		ToUsersID:   mockMessage.ToUserId,
		Subject:     mockMessage.Subject,
		Message:     mockMessage.Message,
	}

	// Mocks service layer for CreateMessage
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.On("CreateMessage", mockResponse).Return(mockResponse, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Basically turns mockMessage into JSON
	b, _ := json.Marshal(mockMessage)

	// Mocks the actual request
	req := httptest.NewRequest("POST", "/messages/", bytes.NewBuffer(b))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.CreateMessage(c)

	// We did indeed CreateMessage :)
	mockMessagesService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Send invalid JSON, expect failure.
func TestMessagesController_CreateMessageInvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocks service layer for CreateMessage
	mockMessagesService := new(mocks.MessagesService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	invalidJson := []byte(`"invalid json"`)

	// Mocks the actual request
	req := httptest.NewRequest("POST", "/messages/", bytes.NewBuffer(invalidJson))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.CreateMessage(c)

	// We did indeed fufill expectations :)
	mockMessagesService.AssertExpectations(t)

	// Achieved failure (intentionally)! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Get an error from the service layer, expect failure.
func TestMessagesController_CreateMessageServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mockMessageStruct struct {
		FromUserId uint
		ToUserId   uint
		Subject    string
		Message    string
	}

	// The mock message we send & expect as JSON
	mockMessage := mockMessageStruct{
		FromUserId: 1,
		ToUserId:   2,
		Subject:    "...hello, there.",
		Message:    "...gEneRAl kEnoBi...",
	}

	mockResponse := models.Messages{
		FromUsersID: mockMessage.FromUserId,
		ToUsersID:   mockMessage.ToUserId,
		Subject:     mockMessage.Subject,
		Message:     mockMessage.Message,
	}

	// Mocks service layer for CreateMessage
	mockMessagesService := new(mocks.MessagesService)

	// Maybe from/to user doesn't exist
	mockMessagesService.On("CreateMessage", mockResponse).
		Return(mockResponse, errors.New("Something went wrong D:"))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Basically turns mockMessage into JSON
	b, _ := json.Marshal(mockMessage)

	// Mocks the actual request
	req := httptest.NewRequest("POST", "/messages/", bytes.NewBuffer(b))

	// This is ensures the content type is set correctly.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.CreateMessage(c)

	// We did indeed CreateMessage :)
	mockMessagesService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// ListAllMessagesForUser tests

// This test ensures that all Message rows are listed for a specific user.
func TestMessagesController_ListAllMessagesForUserSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock messages we expect
	mockMessage := models.Messages{
		FromUsersID: 1,
		ToUsersID:   2,
		Subject:     "i throw my hands up in the air sometimes",
		Message:     "saying ayyyyo, come on let's go",
	}

	mockMessage2 := models.Messages{
		FromUsersID: 1337,
		ToUsersID:   2,
		Subject:     "that's why i never sleep",
		Message:     "bc sleep is the cousin of death",
	}

	type mockMessageStruct struct {
		UserId uint
	}

	// The mock JSON we send
	mockJson := mockMessageStruct{
		UserId: 2,
	}

	var mockMessagesList = []models.Messages{mockMessage, mockMessage2}

	// Mocks service layer
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.On("ListAllMessagesForUser", mockJson.UserId).Return(mockMessagesList, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Basically turns mockJson into actual JSON
	b, _ := json.Marshal(mockJson)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/messages/", bytes.NewBuffer(b))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.ListAllMessagesForUser(c)

	// We did indeed ListAllMessagesForUser :)
	mockMessagesService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestMessagesController_ListAllMessagesForUserInvalidJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Send invalid json, expect failure
	invalidJson := []byte(`invalid json`)

	// Mocks service layer
	mockMessagesService := new(mocks.MessagesService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/messages/", bytes.NewBuffer(invalidJson))

	// This is ensures the content type is set correctly.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.ListAllMessagesForUser(c)

	// Mocks went as expected :)
	mockMessagesService.AssertExpectations(t)

	// Achieved failure (intentionally)! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestMessagesController_ListAllMessagesForUserServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mockMessageStruct struct {
		UserId uint
	}

	// The mock JSON we send
	mockJson := mockMessageStruct{
		UserId: 2,
	}

	// Mocks service layer
	mockMessagesService := new(mocks.MessagesService)

	// Maybe user doesn't exist
	mockMessagesService.On("ListAllMessagesForUser", mockJson.UserId).
		Return(nil, errors.New("Uh oh spaghetti-o..."))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Basically turns mockJson into actual JSON
	b, _ := json.Marshal(mockJson)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/messages/", bytes.NewBuffer(b))

	// This is ensures the content type is set correctly.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.ListAllMessagesForUser(c)

	// Mocks went as expected :)
	mockMessagesService.AssertExpectations(t)

	// Achieved failure (intentionally)! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// FindMessage Tests

// This test ensures that a proper service call
// returns a success for FindMessage.
func TestMessagesController_FindMessageSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect
	mockMessage := models.Messages{
		FromUsersID: 1,
		ToUsersID:   2,
		Subject:     "KAKAROTTTTT",
		Message:     "THIS SCANNER IS BROKEN. IT SAYS YOUR POWER LEVEL EXCEEDS 8,000!!!!!",
	}

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.
		On("FindMessage", uint(1)).
		Return(mockMessage, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/messages/:id", nil)

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	c.AddParam("id", "1")

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.FindMessage(c)

	// We did indeed FindMessage :)
	mockMessagesService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// This test expects an error when message id param is missing.
func TestMessagesController_FindMessageNoMessageId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/messages/:id", nil)

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.FindMessage(c)

	// We did indeed FindOrgUser :)
	mockMessagesService.AssertExpectations(t)

	// id ain't assigned, fam
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// FindMessage returns an error? Then fail.
func TestMessagesController_FindMessageServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock message we expect
	mockMessage := models.Messages{
		FromUsersID: 30,
		ToUsersID:   12,
		Subject:     "Float like a butterfly...",
		Message:     "Sting like a bee, the hands can't hit what the eyes can't see.",
	}

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.
		On("FindMessage", uint(100)).
		Return(mockMessage, errors.New("woah there, cowboy. something went wrong."))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/messages/:id", nil)

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	c.AddParam("id", "100")

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.FindMessage(c)

	// We did indeed FindMessage :)
	mockMessagesService.AssertExpectations(t)

	// Error on Service function, so fail.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// UpdateMessageReadStatus tests

// This test ensures that a proper service call
// returns a success for UpdateMessageReadStatus.
func TestMessagesController_UpdateMessageReadStatusSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock message we expect
	mockMessage := models.Messages{
		FromUsersID: 313,
		ToUsersID:   702,
		Subject:     "man i'm so sick",
		Message:     "i got ambulances pulling me over 'n stuff",
		Read:        true,
	}

	jsonInput := []byte(`{"Read": true}`)

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.
		On("UpdateMessageReadStatus",
			uint(1234), mockMessage.Read).
		Return(mockMessage, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("PUT", "/messages/:id", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	c.AddParam("id", "1234")

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.UpdateMessageReadStatus(c)

	// We did indeed UpdateMessageReadStatus :)
	mockMessagesService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Check that updating read status fails without a message id.
func TestMessagesController_UpdateMessageReadStatusNoMessageId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jsonInput := []byte(`{"Read": true}`)

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("PUT", "/messages/:id", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.UpdateMessageReadStatus(c)

	// We did indeed call everything as expected :)
	mockMessagesService.AssertExpectations(t)

	// Achieved failure (intentionally)! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Expects failure if a request comes with invalid JSON.
func TestMessagesController_UpdateMessageReadStatusInvalidJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jsonInput := []byte(`uwotm8`)

	// Mocks service layer like a cool cat.
	mockMessagesService := new(mocks.MessagesService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller.
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request.
	req := httptest.NewRequest("PUT", "/messages/:id", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	c.AddParam("id", "1234")

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.UpdateMessageReadStatus(c)

	// We did indeed UpdateMessageReadStatus :)
	mockMessagesService.AssertExpectations(t)

	// Expect failure for invalid JSON.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Checks for failure when service layer returns an error for UpdateMessageReadStatus.
func TestMessagesController_UpdateMessageReadStatusServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock message we expect
	mockMessage := models.Messages{
		FromUsersID: 12,
		ToUsersID:   13,
		Subject:     "We're soarin'...",
		Message:     "Flyin',,,There's not a star in Heaven that we can't reaaach!!!!",
		Read:        true,
	}

	jsonInput := []byte(`{"Read": true}`)

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.
		On("UpdateMessageReadStatus",
			uint(2006), mockMessage.Read).
		Return(mockMessage, errors.New("a sad day to be a member of the service layer."))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("PUT", "/messages/:id", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	c.AddParam("id", "2006")

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.UpdateMessageReadStatus(c)

	// We did indeed UpdateMessageReadStatus :)
	mockMessagesService.AssertExpectations(t)

	// Service layer fails? Then we all fail!!!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// DeleteMessage tests

// This test ensures that a proper service call
// returns a success for DeleteMessage.
func TestMessagesController_DeleteMessageSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.
		On("DeleteMessage", uint(321)).
		Return(nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("DELETE", "/messages/:id", nil)

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	c.AddParam("id", "321")

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.DeleteMessage(c)

	// We set our expectations. We met our expectations.
	mockMessagesService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Failure on DeleteMessage with no message id param.
func TestMessagesController_DeleteMessageNoMessageId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("DELETE", "/messages/:id", nil)

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.DeleteMessage(c)

	// We set our expectations. We met our expectations.
	mockMessagesService.AssertExpectations(t)

	// Expect failure because message id is missing.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Ensures failure if service layer returns an error
// for DeleteMessage.
func TestMessagesController_DeleteMessageServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocks service layer like a cool cat
	mockMessagesService := new(mocks.MessagesService)
	mockMessagesService.
		On("DeleteMessage", uint(456)).
		Return(errors.New("me cwumpets >:("))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("DELETE", "/messages/:id", nil)

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	c.AddParam("id", "456")

	// Make a new controller, passing service mock.
	result := NewMessagesController(mockMessagesService)

	result.DeleteMessage(c)

	// We set our expectations. We met our expectations.
	mockMessagesService.AssertExpectations(t)

	// Service layer returned an error, so we expect failure.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
