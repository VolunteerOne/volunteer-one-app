package repository

import (
	"errors"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(models.Organization) (models.Organization, error)
	GetOrganizations() ([]models.Organization, error)
	GetOrganizationById(string) (models.Organization, error)
	UpdateOrganization(models.Organization) (models.Organization, error)
	DeleteOrganization(models.Organization) error
}

type organizationRepository struct {
	DB *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return organizationRepository{
		DB: db,
	}
}

func (r organizationRepository) CreateOrganization(org models.Organization) (models.Organization, error) {
	result := r.DB.Create(&org)

	if result.Error != nil {
		return models.Organization{}, errors.New("creation Failed")
	}

	return org, nil
}

func (r organizationRepository) GetOrganizations() ([]models.Organization, error) {
	var orgs []models.Organization
	result := r.DB.Find(&orgs)

	if result.Error != nil {
		return []models.Organization{}, errors.New("could not retrieve organizations")
	}

	return orgs, nil
}

func (r organizationRepository) GetOrganizationById(id string) (models.Organization, error) {
	var org models.Organization

	result := r.DB.First(&org, id)

	if result.Error != nil {
		return models.Organization{}, errors.New("could not retrieve organization")
	}

	return org, nil
}

func (r organizationRepository) UpdateOrganization(org models.Organization) (models.Organization, error) {
	result := r.DB.Save(&org)

	if result.Error != nil {
		return models.Organization{}, errors.New("could not update organization")
	}

	return org, nil
}

func (r organizationRepository) DeleteOrganization(org models.Organization) error {
	result := r.DB.Where("ID = ?", org.ID).Delete(&org)
	if result.Error != nil {
		return errors.New("could not delete organization")
	}

	return nil
}
