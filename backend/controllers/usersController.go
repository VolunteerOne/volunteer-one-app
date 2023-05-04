package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type UsersController interface {
	Create(c *gin.Context)
	One(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type usersController struct {
	usersService service.UsersService
}

func NewUsersController(s service.UsersService) UsersController {
	return usersController{
		usersService: s,
	}
}

var usersModel = new(models.Users)

// Create ...
func (controller usersController) Create(c *gin.Context) {
	var err error

	// db := database.GetDatabase()

	// Declare a struct for the desired request body
	var body struct {
		Id       uint
		Handle   string
		Email    string
		Password string
		// birthdate datatypes.Date `gorm: "NOT NULL"`
		Birthdate string
		FirstName string
		LastName  string
		// profilePic mediumblob,
		Interests string
		Verified  uint
	}

	// Bind struct to context and check for error
	err = controller.usersService.Bind(c, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	hash, err := controller.usersService.HashPassword([]byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Create the object in the database
	object := models.Users{
		Handle:    body.Handle,
		Email:     body.Email,
		Password:  string(hash),
		Birthdate: body.Birthdate,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		// ProfilePic: body.profilePic,
		Interests: body.Interests,
		Verified:  body.Verified,
	}

	result, err := controller.usersService.CreateUser(object)

	_ = result

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)
}

func (controller usersController) One(c *gin.Context) {

	// Get the id
	id := c.Param("id")

	// Get object from the database
	var object models.Users
	result, err := controller.usersService.OneUser(id, object)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusOK, result)
}

func (controller usersController) Update(c *gin.Context) {

	// Get the existing object
	id := c.Param("id")
	var object models.Users

	result, err := controller.usersService.OneUser(id, object)

	_ = result

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Get updates from the body
	var body struct {
		Id       uint
		Handle   string
		Email    string
		Password string
		// birthdate datatypes.Date `gorm: "NOT NULL"`
		Birthdate string
		FirstName string
		LastName  string
		// profilePic mediumblob,
		Interests string
		Verified  uint
	}
	if err = controller.usersService.Bind(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	object.Handle = body.Handle
	object.Email = body.Email
	object.Password = body.Password
	object.Birthdate = body.Birthdate
	object.FirstName = body.FirstName
	object.LastName = body.LastName
	// object.ProfilePic = body.profilePic
	object.Interests = body.Interests
	object.Verified = body.Verified

	// Update the object

	result1, err := controller.usersService.UpdateUser(object)

	_ = result1

	if err != nil {
		// result = db.Save(&object)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, object)

}

func (controller usersController) Delete(c *gin.Context) {

	// Get the existing object
	id := c.Param("id")
	var object models.Users
	result, err := controller.usersService.OneUser(id, object)

	_ = result

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	result1, err := controller.usersService.DeleteUser(object)

	_ = result1

	if err != nil {
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
