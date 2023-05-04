package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type MessagesService interface {
	CreateMessage(models.Messages) (models.Messages, error)
	ListAllMessagesForUser(uint) ([]models.Messages, error)
	FindMessage(uint) (models.Messages, error)
	UpdateMessageReadStatus(uint, bool) (models.Messages, error)
	DeleteMessage(uint) error
}

type messagesService struct {
	messagesRepository repository.MessagesRepository
}

// Instantiated in router.go
func NewMessagesService(r repository.MessagesRepository) MessagesService {
	return messagesService{
		messagesRepository: r,
	}
}

func (m messagesService) CreateMessage(orgUser models.Messages) (models.Messages, error) {
	return m.messagesRepository.CreateMessage(orgUser)
}

func (m messagesService) ListAllMessagesForUser(userId uint) ([]models.Messages, error) {
	return m.messagesRepository.ListAllMessagesForUser(userId)
}

func (m messagesService) FindMessage(messageId uint) (models.Messages, error) {
	return m.messagesRepository.FindMessage(messageId)
}

func (m messagesService) UpdateMessageReadStatus(messageId uint, read bool) (models.Messages, error) {
	return m.messagesRepository.UpdateMessageReadStatus(messageId, read)
}

func (m messagesService) DeleteMessage(messageId uint) error {
	return m.messagesRepository.DeleteMessage(messageId)
}
