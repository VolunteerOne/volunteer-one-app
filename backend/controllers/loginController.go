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
	Signup(c *gin.Context)
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

// Signup -> Creates a new user profile in the db
func (l loginController) Signup(c *gin.Context) {
	log.Println("[LoginController] Signing up...")

	// Get all the fields from the request body needed for just signup
	var body struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse body",
		})

		return
	}

	hash, err := l.loginService.HashPassword([]byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.Users{Email: body.Email,
		Password:  string(hash),
		FirstName: body.FirstName,
		LastName:  body.LastName}

	result, err := l.loginService.CreateUser(user)

	_ = result

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Respond -> No information needs to be sent back
	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully created",
	})
}

// Login:
// Gets the email and password as a parameter from the request
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
