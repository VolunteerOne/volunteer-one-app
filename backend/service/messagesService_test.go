package service

import (
	"testing"

	"github.com/VolunteerOne/volunteer-one-app/backend/mocks"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
)

// These tests are all trivial since this layer is mostly
// for abstraction and mocking.

func TestMessagesService_CreateMessage(t *testing.T) {
	message := models.Messages{
		FromUsersID: 1,
		ToUsersID:   2,
		Subject:     "For 'tis the sport to have the engineer...",
		Message:     "Hoist with his own petard",
	}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.MessagesRepository)

	// We expect it to return a message and no error.
	mockRepository.On("CreateMessage", message).Return(message, nil)

	messagesService := NewMessagesService(mockRepository)
	messagesService.CreateMessage(message)

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestMessagesService_ListAllMessagesForUser(t *testing.T) {
	message := models.Messages{
		FromUsersID: 22,
		ToUsersID:   33,
		Subject:     "ERMAHGARD",
		Message:     "IMMA FIRIN' MY LAZARRR",
	}

	message2 := models.Messages{
		FromUsersID: 81,
		ToUsersID:   33,
		Subject:     "owo",
		Message:     "uwu",
	}

	var messageList = []models.Messages{message, message2}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.MessagesRepository)

	// We expect it to return a list of messages and no error.
	mockRepository.On("ListAllMessagesForUser", message.ToUsersID).Return(messageList, nil)

	messagesService := NewMessagesService(mockRepository)
	messagesService.ListAllMessagesForUser(message.ToUsersID)

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestMessagesService_FindMessage(t *testing.T) {
	message := models.Messages{
		FromUsersID: 130,
		ToUsersID:   120,
		Subject:     "RE: My anti-itch ointment",
		Message:     "Hey, your anti-itch ointment is ready for pickup now.",
	}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.MessagesRepository)

	// We expect it to return a message and no error.
	mockRepository.On("FindMessage", uint(18)).Return(message, nil)

	messagesService := NewMessagesService(mockRepository)
	messagesService.FindMessage(uint(18))

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestMessagesService_UpdateMessageReadStatus(t *testing.T) {
	message := models.Messages{
		FromUsersID: 900,
		ToUsersID:   800,
		Subject:     "RE: This building smells like updog...",
		Message:     "What's updog?",
		Read:        true,
	}

	// Mock the Repository layer like it's Halloween.
	mockRepository := new(mocks.MessagesRepository)

	// We expect it to return a message and no error.
	mockRepository.
		On("UpdateMessageReadStatus", uint(19), message.Read).
		Return(message, nil)

	messagesService := NewMessagesService(mockRepository)
	messagesService.UpdateMessageReadStatus(uint(19), message.Read)

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}

func TestMessagesService_DeleteMessage(t *testing.T) {
	// Mock the Repository layer like it's Halloween but a different year.
	mockRepository := new(mocks.MessagesRepository)

	// We expect it to return a message and no error.
	mockRepository.
		On("DeleteMessage", uint(1993)).
		Return(nil)

	messagesService := NewMessagesService(mockRepository)
	messagesService.DeleteMessage(uint(1993))

	// Verify our assertions were correct because we always are.
	mockRepository.AssertExpectations(t)
}
