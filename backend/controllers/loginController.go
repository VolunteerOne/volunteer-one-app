package controllers

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// All Controller methods should be defined in the interface
type LoginController interface {
	Login(c *gin.Context)
	SendEmailForPassReset(c *gin.Context)
	PasswordReset(c *gin.Context)
	VerifyAccessToken(c *gin.Context)
	RefreshToken(c *gin.Context)
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

	// 15 minute expire for accessToken
	accessExpire := jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
	// 30 day expire for refreshToken
	refreshExpire := jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))

	accessToken, refreshToken, err := l.loginService.GenerateJWT(user.ID,
		accessExpire, refreshExpire, os.Getenv("JWT_SECRET"), c)

	if err != nil {
		log.Println(err)
		// json status already set in GenerateJWT
		return
	}

	// Store the refresh token in the Delegations table
	var delegations models.Delegations

	// Save the code
	err = l.loginService.SaveRefreshToken(user.ID, refreshToken, delegations)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to save refresh token to DB",
			"success": false,
		})
		return

	}

	// Send the access/refresh token
	c.JSON(http.StatusOK, gin.H{
		"message":       "Successfully logged in",
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
	resetCodeParsed, err := l.loginService.ParseUUID(resetCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse UUID",
			"success": false,
		})
		return
	}

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
	hash, _ := l.loginService.HashPassword([]byte(newPassword))
	if changePasswordErr := l.loginService.ChangePassword(hash, user); changePasswordErr != nil {
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

func (l loginController) VerifyAccessToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User is authenticated",
		"success": true,
	})
	return
}

func (l loginController) RefreshToken(c *gin.Context) {
	// Accept a refresh token, and return a fresh token if available

	// Get the token off the header
	refreshToken, ok := c.Request.Header["Token"]

	if !ok {
		log.Println(`No "token" field in header`)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": `No "token" field in header`,
			"success": false,
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate it
	token, err := middleware.Validate(refreshToken[0], os.Getenv("JWT_SECRET"))

	if err != nil {
		log.Println("Error: Something went wrong when parsing the token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Something went wrong when parsing the token",
			"success": false,
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check if refresh
		if claims["type"] != "refresh" {
			log.Println("Must provide refresh token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Must provide refresh token",
				"success": false,
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Println("Refresh token is expired")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Refresh token is expired",
				"success": false,
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var delegations models.Delegations

		// Get refresh token from DB
		delegations, err = l.loginService.FindRefreshToken(claims["sub"].(float64), delegations)

		// Check that they match -> invalidate the token if so and require user to reauthenticate
		if delegations.RefreshToken != refreshToken[0] {
			log.Println("Invalid refresh token was used - must authenticate again")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid refresh token was used - must authenticate again",
				"success": false,
			})
			c.AbortWithStatus(http.StatusUnauthorized)

			// Delete the refresh token from db -> user will have to reauthenticate
			// User will have access for rest of life of access token but no longer
			l.loginService.DeleteRefreshToken(delegations)

			return
		}

		// They do match so it's valid
		// 15 minute expire for accessToken
		accessExpire := jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
		// 1 day expire for refreshToken
		refreshExpire := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

		accessToken, refreshToken, err := l.loginService.GenerateJWT(delegations.UsersID,
			accessExpire, refreshExpire, os.Getenv("JWT_SECRET"), c)

		if err != nil {
			log.Println(err)
			// json status already set in GenerateJWT
			return
		}

		// Save the code
		err = l.loginService.SaveRefreshToken(delegations.UsersID, refreshToken, delegations)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to save refresh token to DB",
				"success": false,
			})
			return
		}

		// Send the access/refresh token
		c.JSON(http.StatusOK, gin.H{
			"message":       "Token refresh success",
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"success":       true,
		})

	} else {
		log.Println("Could not validate refresh token. Authenticate again")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not validate refresh token. Authenticate again",
			"success": false,
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
