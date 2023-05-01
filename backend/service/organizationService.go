package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/gin-gonic/gin"
)

type OrganizationService interface {
	CreateOrganization(models.Organization) (models.Organization, error)
	GetOrganizations() ([]models.Organization, error)
	GetOrganizationById(string) (models.Organization, error)
	UpdateOrganization(models.Organization) (models.Organization, error)
	DeleteOrganization(models.Organization) error
	Bind(*gin.Context, any) error
}

type organizationService struct {
	organizationRepository repository.OrganizationRepository
}

// CreateOrganization implements OrganizationService
func (s organizationService) CreateOrganization(org models.Organization) (models.Organization, error) {
	return s.organizationRepository.CreateOrganization(org)
}

// DeleteOrganization implements OrganizationService
func (s organizationService) DeleteOrganization(org models.Organization) error {
	return s.organizationRepository.DeleteOrganization(org)
}

// GetOrganizationById implements OrganizationService
func (s organizationService) GetOrganizationById(id string) (models.Organization, error) {
	return s.organizationRepository.GetOrganizationById(id)
}

// GetOrganizations implements OrganizationService
func (s organizationService) GetOrganizations() ([]models.Organization, error) {
	return s.organizationRepository.GetOrganizations()
}

// UpdateOrganization implements OrganizationService
func (s organizationService) UpdateOrganization(org models.Organization) (models.Organization, error) {
	return s.organizationRepository.UpdateOrganization(org)
}

func NewOrganizationService(r repository.OrganizationRepository) OrganizationService {
	return organizationService{
		organizationRepository: r,
	}
}

func (s organizationService) Bind(c *gin.Context, obj any) error {
	return c.Bind(obj)
}
