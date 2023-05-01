package controllers

import (
	"net/http"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type EventMemberController interface {
	Create(c *gin.Context)
	All(c *gin.Context)
	One(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type eventMemberController struct {
	Service service.EventMemberService
}

// All implements EventMemberController
func (controller eventMemberController) All(c *gin.Context) {

	members, err := controller.Service.GetEventMembers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	c.JSON(http.StatusOK, members)
}

// Create implements EventMemberController
func (controller eventMemberController) Create(c *gin.Context) {

	var member models.EventMember
	
	err := c.Bind(&member)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	member, err = controller.Service.CreateEventMember(member)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	c.JSON(http.StatusOK, member)
}

// Delete implements EventMemberController
func (controller eventMemberController) Delete(c *gin.Context) {
	
	id := c.Param("id")

	member, err := controller.Service.GetEventMemberById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = controller.Service.DeleteEventMember(member)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "object deleted successfully",
	})

}

// One implements EventMemberController
func (controller eventMemberController) One(c *gin.Context) {
	id := c.Param("id")

	res, err := controller.Service.GetEventMemberById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, res)
}

// Update implements EventMemberController
func (controller eventMemberController) Update(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		UserId uint
		EventId uint
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	member, err := controller.Service.GetEventMemberById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	member.UserId = body.UserId
	member.EventId = body.EventId

	res, err := controller.Service.UpdateEventMember(member)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, res)
}

func NewEventMemberController(s service.EventMemberService) EventMemberController {
	return eventMemberController{
		Service: s,
	}
}
