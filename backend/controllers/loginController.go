package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (controller LoginController) Login(c *gin.Context) {

	db := database.GetDatabase()
	userInputU := c.Param("email")
	userInputP := c.Param("password")
	// Check if the email exists in the database
	var user models.User
	result := db.Where("email = ?", userInputU).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "email does not exist",
			"success": false,
		})

		return
	}

	// Check if the password matches
	if user.Password != userInputP {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Password does not match",
			"success": false,
		})

		return
	}

	// Respond with true if the username and password match
	c.JSON(http.StatusOK, gin.H{
		"message": "email and password match",
		"success": true,
	})
}
