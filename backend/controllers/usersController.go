package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}

var usersModel = new(models.Users)

// Create ...
func (controller UsersController) Create(c *gin.Context) {
	var err error

	db := database.GetDatabase()

	// Declare a struct for the desired request body
	var body struct {
		Id       uint
		Handle   string
		Email    string
		Password string
		// birthdate datatypes.Date `gorm: "NOT NULL"`
		Birthdate string
		FirstName string
		LastName  string
		// profilePic mediumblob,
		Interests string
		Verified  uint
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
		Handle:    body.Handle,
		Email:     body.Email,
		Password:  body.Password,
		Birthdate: body.Birthdate,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		// ProfilePic: body.profilePic,
		Interests: body.Interests,
		Verified:  body.Verified,
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

func (controller UsersController) All(c *gin.Context) {
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

func (controller UsersController) One(c *gin.Context) {
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

func (controller UsersController) Update(c *gin.Context) {

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
		Id       uint
		Handle   string
		Email    string
		Password string
		// birthdate datatypes.Date `gorm: "NOT NULL"`
		Birthdate string
		FirstName string
		LastName  string
		// profilePic mediumblob,
		Interests string
		Verified  uint
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	object.Handle = body.Handle
	object.Email = body.Email
	object.Password = body.Password
	object.Birthdate = body.Birthdate
	object.FirstName = body.FirstName
	object.LastName = body.LastName
	// object.ProfilePic = body.profilePic
	object.Interests = body.Interests
	object.Verified = body.Verified

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

func (controller UsersController) Delete(c *gin.Context) {
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
