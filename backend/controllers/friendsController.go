package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
)

type friendController struct{}

var friendModel = new(models.Friend)

func (controller friendController) Create(c *gin.Context) {
	var err error
	db := database.GetDatabase()
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

	result := db.Create(&object)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller friendController) Reject(c *gin.Context) {
	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var object models.Friend
	result := db.Find(&object, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	result = db.Delete(&object)
	if result.Error != nil {
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

func (controller friendController) Update(c *gin.Context) {
	db := database.GetDatabase()
	id := c.Param("id")
	var friendDB models.Friend
	result := db.Find(&friendDB, id)

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

	object := models.Friend{
		FriendOneHandle: body.friendOneHandle,
		FriendTwoHandle: body.friendTwoHandle,
		RelationshipBit: "friends",
	}

	results := db.Save(&object)

	if results.Error != nil {

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
	var friend models.Friend
	result := db.First(&friend, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusAccepted, friend)
}

func (controller friendController) All(c *gin.Context) {
	// var err error
	db := database.GetDatabase()
	var friends []models.Friend

	// Get objects from database
	result := db.Find(&friends)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, friends)

}
