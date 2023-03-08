package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
)

type usersController struct{}

var userstModel = new(models.Users)

// Create ...
func (controller usersController) Create(c *gin.Context) {
	var err error
	db := database.GetDatabase()

	// Declare a struct for the desired request body
	var body struct {
		id       uint
		handle   string
		email    string
		password string
		// birthdate datatypes.Date `gorm: "NOT NULL"`
		birthdate string
		firstName string
		lastName  string
		// profilePic mediumblob,
		interests string
		verified  uint
	}

	// Bind struct to context and check for error
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	// Create the object in the database
	object := models.Users{
		Handle:    body.handle,
		Email:     body.email,
		Password:  body.password,
		Birthdate: body.birthdate,
		FirstName: body.firstName,
		LastName:  body.lastName,
		// ProfilePic: body.profilePic,
		Interests: body.interests,
		Verified:  body.verified,
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

func (controller usersController) All(c *gin.Context) {
	// var err error
	db := database.GetDatabase()
	var objects []models.Users

	// Get objects from database
	result := db.Find(&objects)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, objects)

}

func (controller usersController) One(c *gin.Context) {
	db := database.GetDatabase()

	// Get the id
	id := c.Param("id")

	// Get object from the database
	var object models.Users
	result := db.First(&object, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusAccepted, object)
}

func (controller usersController) Update(c *gin.Context) {

	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var object models.Users
	result := db.Find(&object, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Get updates from the body
	var body struct {
		id       uint
		handle   string
		email    string
		password string
		// birthdate datatypes.Date `gorm: "NOT NULL"`
		birthdate string
		firstName string
		lastName  string
		// profilePic mediumblob,
		interests string
		verified  uint
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	object.Handle = body.handle
	object.Email = body.email
	object.Password = body.password
	object.Birthdate = body.birthdate
	object.FirstName = body.firstName
	object.LastName = body.lastName
	// object.ProfilePic = body.profilePic
	object.Interests = body.interests
	object.Verified = body.verified

	// Update the object
	result = db.Save(&object)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller usersController) Delete(c *gin.Context) {
	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var object models.Users
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
