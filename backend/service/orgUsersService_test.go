package service

import (
	"testing"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
)

// These tests are all trivial since this layer is mostly
// for abstraction and mocking.

func TestOrgUsersService_CreateOrgUser(t *testing.T) {
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.OrgUsersRepository)

	// We expect it to return an orgUser and no error.
	mockRepository.On("CreateOrgUser", orgUser).Return(orgUser, nil)

	orgUsersService := NewOrgUsersService(mockRepository)
	orgUsersService.CreateOrgUser(orgUser)

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestOrgUsersService_ListAllOrgUsers(t *testing.T) {
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	orgUser2 := models.OrgUsers{
		UsersID:        3,
		OrganizationID: 4,
	}

	var orgUserList = []models.OrgUsers{orgUser, orgUser2}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.OrgUsersRepository)

	// We expect it to return a list of orgUsers and no error.
	mockRepository.On("ListAllOrgUsers").Return(orgUserList, nil)

	orgUsersService := NewOrgUsersService(mockRepository)
	orgUsersService.ListAllOrgUsers()

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestOrgUsersService_FindOrgUser(t *testing.T) {
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.OrgUsersRepository)

	// We expect it to return an orgUser and no error.
	mockRepository.On("FindOrgUser", orgUser.UsersID, orgUser.OrganizationID).Return(orgUser, nil)

	orgUsersService := NewOrgUsersService(mockRepository)
	orgUsersService.FindOrgUser(orgUser.UsersID, orgUser.OrganizationID)

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestOrgUsersService_UpdateOrgUser(t *testing.T) {
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
		Role:           3,
	}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.OrgUsersRepository)

	// We expect it to return an orgUser and no error.
	mockRepository.
		On("UpdateOrgUser",
			orgUser.UsersID,
			orgUser.OrganizationID,
			orgUser.Role).
		Return(orgUser, nil)

	orgUsersService := NewOrgUsersService(mockRepository)
	orgUsersService.UpdateOrgUser(orgUser.UsersID, orgUser.OrganizationID, orgUser.Role)

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestOrgUsersService_DeleteOrgUser(t *testing.T) {
	orgUser := models.OrgUsers{
		UsersID:        1,
		OrganizationID: 2,
	}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.OrgUsersRepository)

	// We expect it to return an orgUser and no error.
	mockRepository.
		On("DeleteOrgUser",
			orgUser.UsersID,
			orgUser.OrganizationID).
		Return(nil)

	orgUsersService := NewOrgUsersService(mockRepository)
	orgUsersService.DeleteOrgUser(orgUser.UsersID, orgUser.OrganizationID)

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}
