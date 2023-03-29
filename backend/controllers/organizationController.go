package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct{}

var organizationModel = new(models.Organization)

//Create ...
func (controller OrganizationController) Create(c *gin.Context) {
	var err error
	db := database.GetDatabase()

	// Declare a struct for the desired request body
	var body struct {
		Name string
		Description string 
		Verified bool 
		Interests string 
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
	object := models.Organization{
		Name:  body.Name,
		Description: body.Description,
		Verified: body.Verified,
		Interests: body.Interests,
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

func (controller OrganizationController) All(c *gin.Context) {
	// var err error
	db := database.GetDatabase()
	var orgs []models.Organization

	// Get objects from database
	result := db.Find(&orgs)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, orgs)

}

func (controller OrganizationController) One(c *gin.Context) {
	db := database.GetDatabase()

	// Get the id
	id := c.Param("id")

	// Get object from the database
	var org models.Organization
	result := db.First(&org, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusAccepted, org)
}

func (controller OrganizationController) Update(c *gin.Context) {

	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var org models.Organization
	result := db.Find(&org, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Get updates from the body
	var body struct {
		Name string
		Description string 
		Verified bool 
		Interests string 
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	org.Name = body.Name
	org.Description = body.Description
	org.Verified = body.Verified
	org.Interests = body.Interests

	// Update the object
	result = db.Save(&org)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, org)
}

func (controller OrganizationController) Delete(c *gin.Context) {
	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var org models.Organization
	result := db.Find(&org, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	result = db.Delete(&org)
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
