package controllers

import (
	"log"
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
)

// All Controller methods should be defined in the interface
type LoginController interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	PasswordReset(c *gin.Context)
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

func (l loginController) PasswordReset(c *gin.Context) {
	log.Println("Entering PasswordReset send code to email")

	//First find if the email exist
	//if it does then send reset code
	//if not, send back request that it does not exist
	userInputU := c.Param("email")
	var user models.Users
	user, err := l.loginService.FindUserFromEmail(userInputU, user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Email does not exist",
			"success": false,
		})
		return
	}

	//Generate reset code
	resetCode := uuid.New()

	err = l.loginService.SaveResetCodeToUser(resetCode, user)

	// Send reset code to user's email address
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "volunteeronenoreply")
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Password Reset Code")
	mailer.SetBody("text/plain", "Your password reset code is "+resetCode.String())
	if err := gomail.NewDialer("smtp.sendgrid.net", 465, "apikey", "EXAMPLEAPIKEY").DialAndSend(mailer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":        "Failed to send email",
			"success":        false,
			"error messsage": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Email has been sent!",
		"success":   true,
		"resetCode": resetCode,
	})
	return
}
