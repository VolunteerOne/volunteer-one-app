package repository

import (
	"errors"

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
func (r eventRepository) CreateEvent(event models.Event) (models.Event, error) {
	result := r.DB.Create(&event)
	r.DB.Preload("Organization").Find(&event)

	if result.Error != nil {
		return models.Event{}, errors.New("creation failed")
	}

	return event, nil
}

// DeleteEvent implements EventRepository
func (r eventRepository) DeleteEvent(event models.Event) error {
	result := r.DB.Where("ID = ?", event.ID).Delete(&event)
	if result.Error != nil {
		return errors.New("deletion failed")
	}

	return nil
}

// GetEventById implements EventRepository
func (r eventRepository) GetEventById(id string) (models.Event, error) {
	var event models.Event

	result := r.DB.First(&event, id)
	r.DB.Preload("Organization").Find(&event)

	if result.Error != nil {
		return models.Event{}, errors.New("get failed")
	}

	return event, nil
}

// GetEvents implements EventRepository
func (r eventRepository) GetEvents() ([]models.Event, error) {
	var events []models.Event
	result := r.DB.Find(&events)
	r.DB.Preload("Organization").Find(&events)

	if result.Error != nil {
		return []models.Event{}, errors.New("get failed")
	}

	return events, nil
}

// UpdateEvent implements EventRepository
func (r eventRepository) UpdateEvent(event models.Event) (models.Event, error) {
	result := r.DB.Save(&event)
	r.DB.Preload("Organization").Find(&event)

	if result.Error != nil {
		return models.Event{}, errors.New("update failed")
	}

	return event, nil
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return eventRepository{
		DB: db,
	}
}
