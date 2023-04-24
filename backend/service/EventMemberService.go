package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type EventMemberService interface {
	GetEventMembers() ([]models.EventMember, error)
	GetEventMemberById(string) (models.EventMember, error)
	CreateEventMember(models.EventMember) (models.EventMember, error)
	UpdateEventMember(models.EventMember) (models.EventMember, error)
	DeleteEventMember(models.EventMember) error
}

type eventMemberService struct {
	Repository repository.EventMemberRepository
}

// CreateEventMember implements EventMemberService
func (s eventMemberService) CreateEventMember(member models.EventMember) (models.EventMember, error) {
	return s.Repository.CreateEventMember(member)
}

// DeleteEventMember implements EventMemberService
func (s eventMemberService) DeleteEventMember(member models.EventMember) error {
	return s.Repository.DeleteEventMember(member)
}

// GerEventMemberById implements EventMemberService
func (s eventMemberService) GetEventMemberById(id string) (models.EventMember, error) {
	return s.Repository.GetEventMemberById(id)
}

// GetEventMembers implements EventMemberService
func (s eventMemberService) GetEventMembers() ([]models.EventMember, error) {
	return s.Repository.GetEventMembers()
}

// UpdateEventMember implements EventMemberService
func (s eventMemberService) UpdateEventMember(member models.EventMember) (models.EventMember, error) {
	return s.Repository.UpdateEventMember(member)
}

func NewEventMemberService(r repository.EventMemberRepository) EventMemberService {
	return eventMemberService{
		Repository: r,
	}
}
