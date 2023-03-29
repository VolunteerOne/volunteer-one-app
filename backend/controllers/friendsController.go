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
		idOne uint
		idTwo uint
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})
		return
	}

	object := models.Friend{
		friendOneID: body.idOne,
		friendTwoID: body.idTwo,
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



func (controller friendController) Delete(c *gin.Context) {
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
