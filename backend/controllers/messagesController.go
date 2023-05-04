package controllers

import (
	"net/http"
	"strconv"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type MessagesController interface {
	CreateMessage(c *gin.Context)
	ListAllMessagesForUser(c *gin.Context)
	FindMessage(c *gin.Context)
	UpdateMessageReadStatus(c *gin.Context)
	DeleteMessage(c *gin.Context)
}

type messagesController struct {
	messagesService service.MessagesService
}

// Returns the messages controller instantiated in the Router
func NewMessagesController(serv service.MessagesService) MessagesController {
	return messagesController{
		messagesService: serv,
	}
}

// Create a new message from user to user
func (m messagesController) CreateMessage(c *gin.Context) {
	var err error

	// Declare a struct for the desired request body
	var body struct {
		FromUserId uint
		ToUserId   uint
		Subject    string
		Message    string
	}

	// Bind struct to context and check for error
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	// Create message object model, send to next layer
	message := models.Messages{
		FromUsersID: body.FromUserId,
		ToUsersID:   body.ToUserId,
		Subject:     body.Subject,
		Message:     body.Message,
	}

	result, err := m.messagesService.CreateMessage(message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create new Message.",
		})

		return
	}

	// Respond with success
	c.JSON(http.StatusOK, result)
}

// Lists all messages for a specific user (like an inbox)
func (m messagesController) ListAllMessagesForUser(c *gin.Context) {
	// Declare a struct for the desired request body
	var body struct {
		UserId uint
	}

	// Bind struct to context and check for error
	err := c.BindJSON(&body)
  
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	result, err := m.messagesService.ListAllMessagesForUser(body.UserId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve messages for user.",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, result)
}

// Find a message by message ID
func (m messagesController) FindMessage(c *gin.Context) {
	// Get the message ID
	messageId64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id field must be an unsigned integer.",
		})

		return
	}

	// Get object from the database
	messageId := uint(messageId64)
	result, err := m.messagesService.FindMessage(messageId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find message with that id.",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusOK, result)
}

func (m messagesController) UpdateMessageReadStatus(c *gin.Context) {
	// Get the message id
	messageId64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id field must be an unsigned integer.",
		})

		return
	}

	// Get updates from the body
	var body struct {
		Read bool
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	messageId := uint(messageId64)
	result, err := m.messagesService.UpdateMessageReadStatus(messageId, body.Read)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update read status for that message id.",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, result)
}

func (m messagesController) DeleteMessage(c *gin.Context) {
	// Get the message
	messageId64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId field must be an unsigned integer.",
		})

		return
	}

	messageId := uint(messageId64)
	err = m.messagesService.DeleteMessage(messageId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not delete message with that id.",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Message deleted.",
	})
}
