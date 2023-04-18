package repository

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(models.Event) (models.Event, error)
	GetEvents() ([]models.Event, error)
	GetEventById(string) (models.Event, error)
	UpdateEvent(models.Event) (models.Event, error)
	DeleteEvent(models.Event) error
}

type eventRepository struct {
	DB *gorm.DB
}

// CreateEvent implements EventRepository
func (eventRepository) CreateEvent(event models.Event) (models.Event, error) {
	panic("unimplemented")
}

// DeleteEvent implements EventRepository
func (eventRepository) DeleteEvent(event models.Event) error {
	panic("unimplemented")
}

// GetEventById implements EventRepository
func (eventRepository) GetEventById(id string) (models.Event, error) {
	panic("unimplemented")
}

// GetEvents implements EventRepository
func (eventRepository) GetEvents() ([]models.Event, error) {
	panic("unimplemented")
}

// UpdateEvent implements EventRepository
func (eventRepository) UpdateEvent(event models.Event) (models.Event, error) {
	panic("unimplemented")
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return eventRepository{
		DB: db,
	}
}
