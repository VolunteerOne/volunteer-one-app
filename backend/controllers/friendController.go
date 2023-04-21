package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type FriendController interface {
	Create(c *gin.Context)
	Reject(c *gin.Context)
	Accept(c *gin.Context)
	One(c *gin.Context)
	All(c *gin.Context)
}

type friendController struct {
	friendService service.FriendService
}

func NewFriendController(s service.FriendService) FriendController {
	return friendController{
		friendService: s,
	}
}

var friendModel = new(models.Friend)

func (controller friendController) Create(c *gin.Context) {
	var err error
	// db := database.GetDatabase()
	var body struct {
		friendOneHandle string
		friendTwoHandle string
		relationshipBit string
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}

	object := models.Friend{
		FriendOneHandle: body.friendOneHandle,
		FriendTwoHandle: body.friendTwoHandle,
		RelationshipBit: "pending",
	}

	result, err := controller.friendService.CreateFriend(object)

	_ = result

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller friendController) Reject(c *gin.Context) {
	// db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var object models.Friend
	result, err := controller.friendService.OneFriend(id)
	_ = result

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	err1 := controller.friendService.RejectFriend(object)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not delete object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Object deleted successfully",
	})

}

func (controller friendController) Accept(c *gin.Context) {
	db := database.GetDatabase()
	id := c.Param("id")
	var object models.Friend
	result := db.First(&object, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	var body struct {
		friendOneHandle string
		friendTwoHandle string
		relationshipBit string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	object.FriendOneHandle = body.friendOneHandle
	object.FriendTwoHandle = body.friendTwoHandle
	object.RelationshipBit = "friends"

	results1, err := controller.friendService.AcceptFriend(object)
	_ = results1

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller friendController) One(c *gin.Context) {
	db := database.GetDatabase()

	// Get the id
	id := c.Param("id")

	// Get object from the database
	var object models.Friend
	result := db.First(&object, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusAccepted, result)
}

func (controller friendController) All(c *gin.Context) {

	// Get object from the database
	friends, err := controller.friendService.GetFriends()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, friends)

}
