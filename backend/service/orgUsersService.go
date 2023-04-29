package service

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type OrgUsersService interface {
	CreateOrgUser(models.OrgUsers) (models.OrgUsers, error)
	ListAllOrgUsers() ([]models.OrgUsers, error)
	FindOrgUser(uint, uint) (models.OrgUsers, error)
	UpdateOrgUser(uint, uint, uint) (models.OrgUsers, error)
	DeleteOrgUser(uint, uint) error
}

type orgUsersService struct {
	orgUsersRepository repository.OrgUsersRepository
}

// Instantiated in router.go
func NewOrgUsersService(r repository.OrgUsersRepository) OrgUsersService {
	return orgUsersService{
		orgUsersRepository: r,
	}
}

func (o orgUsersService) CreateOrgUser(orgUser models.OrgUsers) (models.OrgUsers, error) {
	return o.orgUsersRepository.CreateOrgUser(orgUser)
}

func (o orgUsersService) ListAllOrgUsers() ([]models.OrgUsers, error) {
	return o.orgUsersRepository.ListAllOrgUsers()
}

func (o orgUsersService) FindOrgUser(userId uint, orgId uint) (models.OrgUsers, error) {
	return o.orgUsersRepository.FindOrgUser(userId, orgId)
}

func (o orgUsersService) UpdateOrgUser(userId uint, orgId uint, role uint) (models.OrgUsers, error) {
	return o.orgUsersRepository.UpdateOrgUser(userId, orgId, role)
}

func (o orgUsersService) DeleteOrgUser(userId uint, orgId uint) error {
	return o.orgUsersRepository.DeleteOrgUser(userId, orgId)
}
