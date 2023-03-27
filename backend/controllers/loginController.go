package controllers

import (
	"log"
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

// All Controller methods should be defined in the interface
type LoginController interface {
	Login(c *gin.Context)
}

// The struct holds the reference to the corresponding service
type loginController struct {
	loginService service.LoginService
}

// Returns the new user controller -> instantiated in router.go
func NewLoginController(s service.LoginService) LoginController {
	return loginController{
		loginService: s,
	}
}

func (l loginController) Login(c *gin.Context) {
	log.Println("[LoginController] Logging in...")

	userInputU := c.Param("email")
	userInputP := c.Param("password")

	var user models.Users

	user, err := l.loginService.FindUserFromEmail(userInputU, user)

	// Email couldn't be found
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":   "Email does not exist",
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
