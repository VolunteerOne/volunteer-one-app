package repository

import (
	"log"
	"strconv"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type MessagesRepository interface {
	CreateMessage(models.Messages) (models.Messages, error)
	ListAllMessagesForUser(uint) ([]models.Messages, error)
	FindMessage(uint) (models.Messages, error)
	UpdateMessageReadStatus(uint, bool) (models.Messages, error)
	DeleteMessage(uint) error
}

type messagesRepository struct {
	DB *gorm.DB
}

// Instantiated in router.go
func NewMessagesRepository(db *gorm.DB) MessagesRepository {
	return messagesRepository{
		DB: db,
	}
}

// Generates new message from user to user
func (m messagesRepository) CreateMessage(messages models.Messages) (models.Messages, error) {
	log.Println("[messagesRepository] Creating message entry...")

	err := m.DB.Create(&messages).Error
	m.DB.Preload("UsersFrom").Find(&messages)
	m.DB.Preload("UsersTo").Find(&messages)

	return messages, err
}

// Lists all messages sent to a specific user
func (m messagesRepository) ListAllMessagesForUser(userId uint) ([]models.Messages, error) {
	log.Println("[messagesRepository] Listing all message rows for user...")

	var messages []models.Messages

	err := m.DB.Find(&messages).Where("ToUserId = ?", userId).Error

	m.DB.Preload("UsersFrom").Find(&messages)
	m.DB.Preload("UsersTo").Find(&messages)

	return messages, err
}

// Finds a message given a specific message ID
func (m messagesRepository) FindMessage(messageId uint) (models.Messages, error) {
	messageIdStr := strconv.FormatUint(uint64(messageId), 10)

	log.Println("[messagesRepository] Finding message with ID (" +
		messageIdStr + ")...")

	var message models.Messages

	err := m.DB.Where("id = ?", messageId).First(&message).Error

	m.DB.Preload("UsersFrom").Find(&message)
	m.DB.Preload("UsersTo").Find(&message)

	return message, err
}

// Updates read status of existing message
func (m messagesRepository) UpdateMessageReadStatus(messageId uint, read bool) (models.Messages, error) {
	messageIdStr := strconv.FormatUint(uint64(messageId), 10)

	log.Println("[messagesRepository] Updating message entry for ID (" +
		messageIdStr + ")...")

	var message models.Messages

	err := m.DB.Where("id = ?", messageId).Find(&message).Error

	if err != nil {
		return message, err
	}

	message.Read = read

	err = m.DB.Save(&message).Error

	m.DB.Preload("UsersFrom").Find(&message)
	m.DB.Preload("UsersTo").Find(&message)

	return message, err
}

// Delete existing message by ID
func (m messagesRepository) DeleteMessage(messageId uint) error {
	messageIdStr := strconv.FormatUint(uint64(messageId), 10)

	log.Println("[messagesRepository] Deleting message for message ID (" +
		messageIdStr + ")...")

	var message models.Messages
	err := m.DB.Find(&message).Error

	if err != nil {
		return err
	}

	err = m.DB.Delete(&message).Error

	return err
}
