package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// *****************************************************
// Tests everything about OrgUsersController :)
// *****************************************************

// CreateOrgUser tests

// This test ensures that a proper service call
// returns a success for a user that was created.
func TestOrgUsersController_CreateOrgUserSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we send & expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 1,
	}

	// Mocks service layer for CreateOrgUser
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.On("CreateOrgUser", mockOrgUser).Return(mockOrgUser, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Basically turns mockOrgUser into JSON
	b, _ := json.Marshal(mockOrgUser)

	// Mocks the actual request
	req := httptest.NewRequest("POST", "/orgUsers/", bytes.NewBuffer(b))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.CreateOrgUser(c)

	// We did indeed CreateOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Send invalid JSON, expect failure.
func TestOrgUsersController_CreateOrgUserInvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocks service layer for CreateOrgUser
	mockOrgUsersService := new(mocks.OrgUsersService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	invalidJson := []byte(`"invalid json"`)

	// Mocks the actual request
	req := httptest.NewRequest("POST", "/orgUsers/", bytes.NewBuffer(invalidJson))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.CreateOrgUser(c)

	// We did indeed fufill expectations :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved failure (intentionally)! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Get an error from the service layer, expect failure.
func TestOrgUsersController_CreateOrgUserServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we send & expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 1,
	}

	// Mocks service layer for CreateOrgUser
	mockOrgUsersService := new(mocks.OrgUsersService)

	// Maybe user/org doesn't exist
	mockOrgUsersService.On("CreateOrgUser", mockOrgUser).
		Return(mockOrgUser, errors.New("Something went wrong D:"))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Basically turns mockOrgUser into JSON
	b, _ := json.Marshal(mockOrgUser)

	// Mocks the actual request
	req := httptest.NewRequest("POST", "/orgUsers/", bytes.NewBuffer(b))

	// This is ensures the content type is set correctly.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.CreateOrgUser(c)

	// We did indeed CreateOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// ListAllOrgUsers tests

// This test ensures that all OrgUser rows are listed.
func TestOrgUsersController_ListAllOrgUsersSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 1,
	}

	mockOrgUser2 := models.OrgUsers{
		UsersID:        2,
		OrganizationID: 3,
	}

	var mockOrgUserList = []models.OrgUsers{mockOrgUser, mockOrgUser2}

	// Mocks service layer
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.On("ListAllOrgUsers").Return(mockOrgUserList, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/orgUsers/", nil)

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.ListAllOrgUsers(c)

	// We did indeed ListAllOrgUsers :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestOrgUsersController_ListAllOrgUsersServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocks service layer
	mockOrgUsersService := new(mocks.OrgUsersService)

	// Maybe user/org doesn't exist
	mockOrgUsersService.On("ListAllOrgUsers").
		Return(nil, errors.New("Uh oh spaghetti-o..."))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/orgUsers/", nil)

	// This is ensures the content type is set correctly.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.ListAllOrgUsers(c)

	// Mocks went as expected :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved failure (intentionally)! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// FindOrgUser Tests

// This test ensures that a proper service call
// returns a success for FindOrgUser.
func TestOrgUsersController_FindOrgUserSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	jsonInput := []byte(`{"OrganizationId": 2}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.
		On("FindOrgUser",
			mockOrgUser.UsersID,
			mockOrgUser.OrganizationID).
		Return(mockOrgUser, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.FindOrgUser(c)

	// We did indeed FindOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// This test expects an error when userId param is missing.
func TestOrgUsersController_FindOrgUserNoUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jsonInput := []byte(`{"OrganizationId": 2}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.FindOrgUser(c)

	// We did indeed FindOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// userId ain't assigned, fam
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// This test sends invalid JSON & expects to fail.
func TestOrgUsersController_FindOrgUserInvalidJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	invalidJson := []byte(`orgwhatidnow`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/orgUsers/:userId", bytes.NewBuffer(invalidJson))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.FindOrgUser(c)

	// We did indeed FindOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Invalid JSON so fail.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// FindOrgUser returns an error? Then fail.
func TestOrgUsersController_FindOrgUserServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	jsonInput := []byte(`{"OrganizationId": 2}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.
		On("FindOrgUser",
			mockOrgUser.UsersID,
			mockOrgUser.OrganizationID).
		Return(mockOrgUser, errors.New("woah there, cowboy. something went wrong."))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("GET", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.FindOrgUser(c)

	// We did indeed FindOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Error on Service function, so fail.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// UpdateOrgUser tests

// This test ensures that a proper service call
// returns a success for UpdateOrgUser.
func TestOrgUsersController_UpdateOrgUserSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           7,
	}

	jsonInput := []byte(`{"OrganizationId": 2, "Role": 7}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.
		On("UpdateOrgUser",
			mockOrgUser.UsersID,
			mockOrgUser.OrganizationID,
			mockOrgUser.Role).
		Return(mockOrgUser, nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("PUT", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.UpdateOrgUser(c)

	// We did indeed UpdateOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Check that userId is omitted, fail.
func TestOrgUsersController_UpdateOrgUserNoUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jsonInput := []byte(`{"OrganizationId": 2, "Role": 7}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("PUT", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.UpdateOrgUser(c)

	// We did indeed call everything as expected :)
	mockOrgUsersService.AssertExpectations(t)

	// Achieved failure (intentionally)! Huzzah!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Expects failure if a request comes with invalid JSON.
func TestOrgUsersController_UpdateOrgUserInvalidJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect:
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           7,
	}

	jsonInput := []byte(`uwotm8`)

	// Mocks service layer like a cool cat.
	mockOrgUsersService := new(mocks.OrgUsersService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller.
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request.
	req := httptest.NewRequest("PUT", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.UpdateOrgUser(c)

	// We did indeed UpdateOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Expect failure for invalid JSON.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Checks for failure when service layer returns an error for UpdateOrgUser.
func TestOrgUsersController_UpdateOrgUserServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we expect
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           7,
	}

	jsonInput := []byte(`{"OrganizationId": 2, "Role": 7}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.
		On("UpdateOrgUser",
			mockOrgUser.UsersID,
			mockOrgUser.OrganizationID,
			mockOrgUser.Role).
		Return(mockOrgUser, errors.New("a sad day to be a member of the service layer."))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("PUT", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.UpdateOrgUser(c)

	// We did indeed UpdateOrgUser :)
	mockOrgUsersService.AssertExpectations(t)

	// Service layer fails? Then we all fail!!!
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// DeleteOrgUser tests

// This test ensures that a proper service call
// returns a success for DeleteOrgUser.
func TestOrgUsersController_DeleteOrgUserSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we send
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	jsonInput := []byte(`{"OrganizationId": 2}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.
		On("DeleteOrgUser",
			mockOrgUser.UsersID,
			mockOrgUser.OrganizationID).
		Return(nil)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("DELETE", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.DeleteOrgUser(c)

	// We set our expectations. We met our expectations.
	mockOrgUsersService.AssertExpectations(t)

	// Achieved success! Huzzah!
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Failure on DeleteOrgUser with no userId param.
func TestOrgUsersController_DeleteOrgUserNoUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonInput := []byte(`{"OrganizationId": 2}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("DELETE", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.DeleteOrgUser(c)

	// We set our expectations. We met our expectations.
	mockOrgUsersService.AssertExpectations(t)

	// Expect failure because userId is missing.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// This test ensures that a proper service call
// returns a success for DeleteOrgUser.
func TestOrgUsersController_DeleteOrgUserInvalidJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we send
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	invalidJson := []byte(`u? and me? u are trying to delete me? asking all them questions.`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("DELETE", "/orgUsers/:userId", bytes.NewBuffer(invalidJson))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.DeleteOrgUser(c)

	// We set our expectations. We met our expectations.
	mockOrgUsersService.AssertExpectations(t)

	// Expect failure on invalid JSON input.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// Ensures failure if service layer returns an error
// for DeleteOrgUser.
func TestOrgUsersController_DeleteOrgUserServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// The mock user we send
	mockOrgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	jsonInput := []byte(`{"OrganizationId": 2}`)

	// Mocks service layer like a cool cat
	mockOrgUsersService := new(mocks.OrgUsersService)
	mockOrgUsersService.
		On("DeleteOrgUser",
			mockOrgUser.UsersID,
			mockOrgUser.OrganizationID).
		Return(errors.New("MEGAN. - drake & josh"))

	rr := httptest.NewRecorder()

	// Pass this instead of a real context to the controller
	c, _ := gin.CreateTestContext(rr)

	// Mocks the actual request
	req := httptest.NewRequest("DELETE", "/orgUsers/:userId", bytes.NewBuffer(jsonInput))

	// This is ensures the content type is set correctly.
	// It's necessary for the test to pass.
	req.Header.Set("Content-Type", "application/json")

	// Add the request to the actual test context.
	c.Request = req
	userIdString := strconv.FormatUint(uint64(mockOrgUser.UsersID), 10)

	c.AddParam("userId", userIdString)

	// Make a new controller, passing service mock.
	result := NewOrgUsersController(mockOrgUsersService)

	result.DeleteOrgUser(c)

	// We set our expectations. We met our expectations.
	mockOrgUsersService.AssertExpectations(t)

	// Service layer returned an error, so we expect failure.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
