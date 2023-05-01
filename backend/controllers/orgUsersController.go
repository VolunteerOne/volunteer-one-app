package controllers

import (
	"net/http"
	"strconv"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type OrgUsersController interface {
	CreateOrgUser(c *gin.Context)
	ListAllOrgUsers(c *gin.Context)
	FindOrgUser(c *gin.Context)
	UpdateOrgUser(c *gin.Context)
	DeleteOrgUser(c *gin.Context)
}

type orgUsersController struct {
	orgUsersService service.OrgUsersService
}

// Returns the org user controller instantiated in the Router
func NewOrgUsersController(serv service.OrgUsersService) OrgUsersController {
	return orgUsersController{
		orgUsersService: serv,
	}
}

// Create new role tied to a user and organization
func (o orgUsersController) CreateOrgUser(c *gin.Context) {
	var err error

	// Declare a struct for the desired request body
	var body struct {
		UsersId        uint
		OrganizationId uint
		Role           uint
	}

	// Bind struct to context and check for error
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	// Create orgUser object model, send to next layer
	orgUser := models.OrgUsers{
		UsersID:        body.UsersId,
		OrganizationID: body.OrganizationId,
		Role:           body.Role,
	}

	result, err := o.orgUsersService.CreateOrgUser(orgUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create new OrgUser.",
		})

		return
	}

	// Respond with success
	c.JSON(http.StatusOK, result)
}

// Lists all members with roles in organizations
func (o orgUsersController) ListAllOrgUsers(c *gin.Context) {
	result, err := o.orgUsersService.ListAllOrgUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve organization users.",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, result)
}

// Find a user's organization roles by User ID
func (o orgUsersController) FindOrgUser(c *gin.Context) {
	// Get the userId
	userId64, err := strconv.ParseUint(c.Param("userId"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId field must be an unsigned integer.",
		})

		return
	}

	var body struct {
		OrganizationId uint
	}

	// Bind struct to context and check for errors
	err = c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	// Get object from the database
	userId := uint(userId64)
	result, err := o.orgUsersService.FindOrgUser(userId, body.OrganizationId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find OrgUser with that userId/orgId.",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusOK, result)
}

func (o orgUsersController) UpdateOrgUser(c *gin.Context) {
	// Get the userId
	userId64, err := strconv.ParseUint(c.Param("userId"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId field must be an unsigned integer.",
		})

		return
	}

	// Get updates from the body
	var body struct {
		OrganizationId uint
		Role           uint
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	userId := uint(userId64)
	result, err := o.orgUsersService.UpdateOrgUser(userId, body.OrganizationId, body.Role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update OrgUser with that userId/orgId.",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, result)
}

func (o orgUsersController) DeleteOrgUser(c *gin.Context) {
	// Get the userId
	userId64, err := strconv.ParseUint(c.Param("userId"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId field must be an unsigned integer.",
		})

		return
	}

	var body struct {
		OrganizationId uint
	}

	// Bind struct to context and check for errors
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid.",
		})

		return
	}

	userId := uint(userId64)
	err = o.orgUsersService.DeleteOrgUser(userId, body.OrganizationId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not delete OrgUser.",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "OrgUser deleted.",
	})

}
