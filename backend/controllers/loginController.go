package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
)

// All Controller methods should be defined in the interface
type LoginController interface {
	Login(c *gin.Context)
	SendEmailForPassReset(c *gin.Context)
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
	// Compare the hashed password with the user input password
	erros := l.loginService.CompareHashedAndUserPass([]byte(user.Password), userInputP)
	if erros != nil {
		// Password does not match
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Password does not match",
			"success": false,
		})
		return
	}

	// Generate the JWT Access token

	// 30 minute accessExpire time
	accessExpire := time.Now().Add(time.Minute * 30)
	accessToken, err := l.loginService.GenerateJWT(user.ID, accessExpire, os.Getenv("JWT_SECRET"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create access token",
			"success": false,
		})
	}

	// Generate the JWT Refresh token
	// TODO: Store in DB

	// 7 day expire time
	refreshExpire := time.Now().Add(time.Hour * 24 * 7)
	refreshToken, err := l.loginService.GenerateJWT(user.ID, refreshExpire, os.Getenv("JWT_SECRET"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create refresh token",
			"success": false,
		})
	}

	// Send the access/refresh token
	c.JSON(http.StatusOK, gin.H{
		"message":       "email and password match",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"success":       true,
	})
}

func (l loginController) SendEmailForPassReset(c *gin.Context) {
	log.Println("Entering SendEmailForPassReset function")

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

	//Send reset code to user's email address
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "edwardsung4217@gmail.com") //need to replace this with proper volunteer email
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Password Reset Code")
	mailer.SetBody("text/plain", "Your password reset code is "+resetCode.String())
	if err := gomail.NewDialer("smtp.sendgrid.net", 465, "apikey", "APIKEY").DialAndSend(mailer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to send email",
			"success": false,
			//"error messsage": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Email has been sent!",
		"success": true,
		//"resetCode": resetCode,
	})
	return

}

func (l loginController) PasswordReset(c *gin.Context) {
	email := c.Param("email")
	resetCode := c.Param("resetcode")
	resetCodeParsed, err := uuid.Parse(resetCode)
	newPassword := c.Param("newpassword")

	var user models.Users

	//Retrieve user's record by their email
	user, ero := l.loginService.FindUserFromEmail(email, user)
	if ero != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Email does not exist",
			"success": false,
		})
		return
	}
	//See if reset code is matched with the one they provided
	if user.ResetCode != resetCodeParsed {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to provde correct reset code",
			"success": false,
		})
		return
	}
	hash, err := l.loginService.HashPassword([]byte(newPassword))
	if changePasswordErr := l.loginService.ChangePassword(hash, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":        "Failed to change password",
			"success":        false,
			"error messsage": changePasswordErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your password has been sucessfully changed!",
		"success": true,
	})
	return
}
