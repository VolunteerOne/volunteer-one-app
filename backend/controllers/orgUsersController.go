package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
)

type OrgUsersController struct{}

// Create organization user
func (controller OrgUsersController) CreateOrgUser(c *gin.Context) {
	var err error
	db := database.GetDatabase()

	// Declare a struct for the desired request body
	var body struct {
		UserId         uint
		OrganizationId uint
		Role           uint
	}

	// Bind struct to context and check for error
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	// Create the object in the database
	orgUser := models.OrgUsers{
		UsersID:        body.UserId,
		OrganizationID: body.OrganizationId,
		Role:           body.Role,
	}

	result := db.Create(&orgUser)

	db.Preload("Users").Find(&orgUser)
	db.Preload("Organization").Find(&orgUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create a new Org User.",
		})

		return
	}

	// Respond with success
	c.JSON(http.StatusOK, orgUser)
}

func (controller OrgUsersController) ListAllOrgUsers(c *gin.Context) {
	// var err error
	db := database.GetDatabase()
	var orgUsers []models.OrgUsers

	// Get objects from database
	result := db.Find(&orgUsers)

	db.Preload("Users").Find(&orgUsers)
	db.Preload("Organization").Find(&orgUsers)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve organization users.",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, orgUsers)
}

func (controller OrgUsersController) FindOrgUser(c *gin.Context) {
	db := database.GetDatabase()

	// Get the id
	id := c.Param("id")

	// Get object from the database
	var orgUsers models.OrgUsers
	result := db.First(&orgUsers, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find Org User with that ID.",
		})

		return
	}

	db.Preload("Users").Find(&orgUsers)
	db.Preload("Organization").Find(&orgUsers)

	// Return the object
	c.JSON(http.StatusAccepted, orgUsers)
}

func (controller OrgUsersController) UpdateOrgUser(c *gin.Context) {
	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var orgUser models.OrgUsers
	result := db.Find(&orgUser, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve an Org User with that ID.",
		})

		return
	}

	// Get updates from the body
	var body struct {
		Role uint
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	orgUser.Role = body.Role

	// Update the object
	result = db.Save(&orgUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update object",
		})

		return
	}

	db.Preload("Users").Find(&orgUser)
	db.Preload("Organization").Find(&orgUser)

	// Respond
	c.JSON(http.StatusOK, orgUser)
}

func (controller OrgUsersController) DeleteOrgUser(c *gin.Context) {
	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var orgUser models.OrgUsers
	result := db.Find(&orgUser, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could find an Org User with that ID.",
		})

		return
	}

	// Delete the object
	result = db.Delete(&orgUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not delete that Org User.",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Org User deleted.",
	})

}
