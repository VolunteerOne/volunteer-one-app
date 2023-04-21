package controllers

import (
	"net/http"
	"time"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

type EventController interface {
	Create(*gin.Context)
	All(*gin.Context)
	One(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type eventController struct {
	eventService service.EventService
}

// All implements EventController
func (controller eventController) All(c *gin.Context) {
	events, err := controller.eventService.GetEvents();

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, events)
}

// Create implements EventController
func (controller eventController) Create(c *gin.Context) {
	var err error

	var body struct {
		OrganizationID  uint
		Name        	string
		Address			string
		Date 			time.Time
		Description 	string
		Interests		string
		Skills			string
		GoodFor			string
		CauseAreas		string
		Requirements 	string	
	}

	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return
	}

	event := models.Event {
		OrganizationID: body.OrganizationID,
		Name: body.Name,
		Address: body.Address,
		Date: body.Date,
		Description: body.Description,
		Interests: body.Interests,
		Skills: body.Skills,
		GoodFor: body.GoodFor,
		CauseAreas: body.CauseAreas,
		Requirements: body.Requirements,
	}

	res, err := controller.eventService.CreateEvent(event);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return
	}
	
	c.JSON(http.StatusOK, res)
}

// Delete implements EventController
func (controller eventController) Delete(c *gin.Context) {
	// Get the existing object
	id := c.Param("id")
	org, err := controller.eventService.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Delete the object
	err = controller.eventService.DeleteEvent(org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Object deleted successfully",
	})}

// One implements EventController
func (controller eventController) One(c *gin.Context) {
	// Get the id
	id := c.Param("id")

	// Get object from the database
	event, err := controller.eventService.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Return the object
	c.JSON(http.StatusAccepted, event)}

// Update implements EventController
func (controller eventController) Update(c *gin.Context) {
	id := c.Param("id")

	event, err := controller.eventService.GetEventById(id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Get updates from the body
	var body struct {
		OrganizationID  uint
		Name        	string
		Address			string
		Date 			time.Time
		Description 	string
		Interests		string
		Skills			string
		GoodFor			string
		CauseAreas		string
		Requirements 	string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	event.OrganizationID = body.OrganizationID
	event.Name = body.Name
	event.Address = body.Address
	event.Date = body.Date
	event.Description = body.Description
	event.Interests = body.Interests
	event.Skills = body.Skills			
	event.GoodFor = body.GoodFor			
	event.CauseAreas = body.CauseAreas		
	event.Requirements = body.Requirements 	

	// Update the object
	result, err := controller.eventService.UpdateEvent(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, result)
}

func NewEventController(s service.EventService) EventController {
	return eventController{
		eventService: s,
	}
}
