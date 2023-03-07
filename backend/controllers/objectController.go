package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
)

type ObjectController struct{}

var objectModel = new(models.Object)

//Create ...
func (controller ObjectController) Create(c *gin.Context) {
	var err error
	db := database.GetDatabase()

	// Declare a struct for the desired request body
	var body struct {
		Name  string
		Value string
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
	object := models.Object{
		Name:  body.Name,
		Value: body.Value,
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

func (controller ObjectController) All(c *gin.Context) {
	// var err error
	db := database.GetDatabase()
	var objects []models.Object

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

func (controller ObjectController) One(c *gin.Context) {
	db := database.GetDatabase()

	// Get the id
	id := c.Param("id")

	// Get object from the database
	var object models.Object
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

func (controller ObjectController) Update(c *gin.Context) {

	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var object models.Object
	result := db.Find(&object, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Get updates from the body
	var body struct {
		Name  string
		Value string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	object.Name = body.Name
	object.Value = body.Value

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

func (controller ObjectController) Delete(c *gin.Context) {
	db := database.GetDatabase()

	// Get the existing object
	id := c.Param("id")
	var object models.Object
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
