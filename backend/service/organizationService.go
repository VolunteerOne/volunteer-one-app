package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/google/uuid"
)

type OrganizationService interface {
	CreateOrganization(models.Organization) (models.Organization, error)
	GetOrganizations() ([]models.Organization, error)
	GetOrganizationById(uuid.UUID) (models.Organization, error)
	UpdateOrganization(uuid.UUID, models.Organization) (models.Organization, error)
	DeleteOrganization(uuid.UUID) error
}

type organizationService struct {
	organizationRepository repository.OrganizationRepository
}

// CreateOrganization implements OrganizationService
func (s organizationService) CreateOrganization(org models.Organization) (models.Organization, error) {
	return s.organizationRepository.CreateOrganization(org)
}

// DeleteOrganization implements OrganizationService
func (s organizationService) DeleteOrganization(id uuid.UUID) error {
	return s.organizationRepository.DeleteOrganization(id)
}

// GetOrganizationById implements OrganizationService
func (s organizationService) GetOrganizationById(id uuid.UUID) (models.Organization, error) {
	return s.organizationRepository.GetOrganizationById(id)
}

// GetOrganizations implements OrganizationService
func (s organizationService) GetOrganizations() ([]models.Organization, error) {
	return s.organizationRepository.GetOrganizations()
}

// UpdateOrganization implements OrganizationService
func (s organizationService) UpdateOrganization(id uuid.UUID, org models.Organization) (models.Organization, error) {
	return s.organizationRepository.UpdateOrganization(id, org)
}

func NewOrganizationService(r repository.OrganizationRepository) OrganizationService {
	return organizationService{
		organizationRepository: r,
	}
}
