package controllers

import (
	"net/http"

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
	var body struct {
		FriendOneHandle string
		FriendTwoHandle string
	}
	err = controller.friendService.Bind(c, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}

	object := models.Friend{
		FriendOneHandle: body.FriendOneHandle,
		FriendTwoHandle: body.FriendTwoHandle,
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

	id := c.Param("id")

	result, err := controller.friendService.OneFriend(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	err1 := controller.friendService.RejectFriend(result)

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
	id := c.Param("id")
	result, err1 := controller.friendService.OneFriend(id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	result.RelationshipBit = "friends"

	results1, err := controller.friendService.AcceptFriend(result)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, results1)
}

func (controller friendController) One(c *gin.Context) {

	id := c.Param("id")

	result, err1 := controller.friendService.OneFriend(id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusOK, result)
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
