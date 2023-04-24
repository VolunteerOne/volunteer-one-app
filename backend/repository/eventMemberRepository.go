package repository

import (
	"errors"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type EventMemberRepository interface {
	GetEventMembers() ([]models.EventMember, error)
	GetEventMemberById(string) (models.EventMember, error)
	CreateEventMember(models.EventMember) (models.EventMember, error)
	UpdateEventMember(models.EventMember) (models.EventMember, error)
	DeleteEventMember(models.EventMember) error
}

type eventMemberRepository struct {
	DB *gorm.DB
}

func (r eventMemberRepository) CreateEventMember(member models.EventMember) (models.EventMember, error) {
	res := r.DB.Create(&member);
	
	r.DB.Preload("Users").Find(&member)
	r.DB.Preload("Event").Find(&member)

	if res.Error != nil {
		return models.EventMember{}, errors.New("creation failed")
	}

	return member, nil
}

func (r eventMemberRepository) DeleteEventMember(member models.EventMember) error {
	res := r.DB.Delete(&member)

	if res.Error != nil {
		return errors.New("delete failed");
	}

	return nil
}

func (r eventMemberRepository) GetEventMemberById(id string) (models.EventMember, error) {
	var member models.EventMember

	res := r.DB.First(&member, id)

	r.DB.Preload("Users").Find(&member)
	r.DB.Preload("Event").Find(&member)

	if res.Error != nil {
		return models.EventMember{}, errors.New("get by id failed")
	}

	return member, nil
}

func (r eventMemberRepository) GetEventMembers() ([]models.EventMember, error) {
	var members []models.EventMember

	res := r.DB.Find(&members)

	r.DB.Preload("Users").Find(&members)
	r.DB.Preload("Event").Find(&members)

	if res.Error != nil {
		return []models.EventMember{}, errors.New("get failed")
	}

	return members, nil
}

func (r eventMemberRepository) UpdateEventMember(member models.EventMember) (models.EventMember, error) {

	res := r.DB.Save(&member)

	r.DB.Preload("Users").Find(&member)
	r.DB.Preload("Event").Find(&member)

	if res.Error != nil {
		return models.EventMember{}, errors.New("update failed")
	}

	return member, nil

}

func NewEventMemberRepository(db *gorm.DB) EventMemberRepository {
	return eventMemberRepository{
		DB: db,
	}
}
