package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type EventService interface {
	CreateEvent(models.Event) (models.Event, error)
	GetEvents() ([]models.Event, error)
	GetEventById(string) (models.Event, error)
	UpdateEvent(models.Event) (models.Event, error)
	DeleteEvent(models.Event) error
}

type eventService struct {
	eventRepository repository.EventRepository
}

// CreateEvent implements EventService
func (s eventService) CreateEvent(event models.Event) (models.Event, error) {
	panic("unimplemented")
}

// DeleteEvent implements EventService
func (s eventService) DeleteEvent(event models.Event) error {
	panic("unimplemented")
}

// GetEventById implements EventService
func (s eventService) GetEventById(id string) (models.Event, error) {
	panic("unimplemented")
}

// GetEvents implements EventService
func (s eventService) GetEvents() ([]models.Event, error) {
	panic("unimplemented")
}

// UpdateEvent implements EventService
func (s eventService) UpdateEvent(event models.Event) (models.Event, error) {
	panic("unimplemented")
}

func NewEventService(r repository.EventRepository) EventService {
	return eventService{
		eventRepository: r,
	}
}
