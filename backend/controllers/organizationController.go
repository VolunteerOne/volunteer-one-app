package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type OrganizationController interface{
	Create(*gin.Context)
	All(*gin.Context)
	One(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type organizationController struct {
	organizationService service.OrganizationService
}

func NewOrganizationController(s service.OrganizationService) OrganizationController{
	return organizationController {
		organizationService: s,
	}
}

// var organizationModel = new(models.Organization)

//Create ...
func (controller organizationController) Create(c *gin.Context) {
	var err error

	// Declare a struct for the desired request body
	var body struct {
		Name string
		Description string 
		Verified bool 
		Interests string 
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
	object := models.Organization{
		Name:  body.Name,
		Description: body.Description,
		Verified: body.Verified,
		Interests: body.Interests,
	}

	res, err := controller.organizationService.CreateOrganization(object)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Creation failed",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, res)
}

func (controller organizationController) All(c *gin.Context) {

	// Get objects from database
	orgs, err := controller.organizationService.GetOrganizations()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve objects",
		})

		return
	}

	// Return the array of objects
	c.JSON(http.StatusOK, orgs)

}

func (controller organizationController) One(c *gin.Context) {
	// Get the id
	id := c.Param("id")

	// Get object from the database
	org, err := controller.organizationService.GetOrganizationById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Return the object
	c.JSON(http.StatusAccepted, org)
}

func (controller organizationController) Update(c *gin.Context) {

	id := c.Param("id")

	org, err := controller.organizationService.GetOrganizationById(id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Get updates from the body
	var body struct {
		Name string
		Description string 
		Verified bool 
		Interests string 
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is invalid",
		})

		return
	}

	org.Name = body.Name
	org.Description = body.Description
	org.Verified = body.Verified
	org.Interests = body.Interests

	// Update the object
	result, err := controller.organizationService.UpdateOrganization(org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update object",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, result)
}

func (controller organizationController) Delete(c *gin.Context) {

	// Get the existing object
	id := c.Param("id")
	org, err := controller.organizationService.GetOrganizationById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not retrieve object",
		})

		return
	}

	// Delete the object
	err = controller.organizationService.DeleteOrganization(org)
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