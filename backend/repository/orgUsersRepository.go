package repository

import (
	"log"
	"strconv"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type OrgUsersRepository interface {
	CreateOrgUser(models.OrgUsers) (models.OrgUsers, error)
	ListAllOrgUsers() ([]models.OrgUsers, error)
	FindOrgUser(uint, uint) (models.OrgUsers, error)
	UpdateOrgUser(uint, uint, uint) (models.OrgUsers, error)
	DeleteOrgUser(uint, uint) error
}

type orgUsersRepository struct {
	DB *gorm.DB
}

// Instantiated in router.go
func NewOrgUsersRepository(db *gorm.DB) OrgUsersRepository {
	return orgUsersRepository{
		DB: db,
	}
}

// Generates new user role for an organization
func (o orgUsersRepository) CreateOrgUser(orgUser models.OrgUsers) (models.OrgUsers, error) {
	log.Println("[orgsUsersRepository] Creating OrgUser entry...")

	err := o.DB.Create(&orgUser).Error
	o.DB.Preload("Users").Find(&orgUser)
	o.DB.Preload("Organization").Find(&orgUser)

	return orgUser, err
}

// Lists all members with roles in organizations
func (o orgUsersRepository) ListAllOrgUsers() ([]models.OrgUsers, error) {
	log.Println("[orgsUsersRepository] Listing all OrgUser rows...")

	var orgUsers []models.OrgUsers

	err := o.DB.Find(&orgUsers).Error

	o.DB.Preload("Users").Find(&orgUsers)
	o.DB.Preload("Organization").Find(&orgUsers)

	return orgUsers, err
}

// Finds a user with a role in an organization by User ID
func (o orgUsersRepository) FindOrgUser(userId uint, orgId uint) (models.OrgUsers, error) {
	userIdStr := strconv.FormatUint(uint64(userId), 10)
	orgIdStr := strconv.FormatUint(uint64(orgId), 10)

	log.Println("[orgsUsersRepository] Finding OrgUser entry with ID (" +
		userIdStr + ") in Org ID (" + orgIdStr + ")...")

	var orgUser models.OrgUsers

	err := o.DB.Where("users_id = ? AND organization_id = ?", userId, orgId).First(&orgUser).Error

	o.DB.Preload("Users").Find(&orgUser)
	o.DB.Preload("Organization").Find(&orgUser)

	return orgUser, err
}

// Updates role for existing OrgUser object by ID
func (o orgUsersRepository) UpdateOrgUser(userId uint, orgId uint, role uint) (models.OrgUsers, error) {
	userIdStr := strconv.FormatUint(uint64(userId), 10)
	orgIdStr := strconv.FormatUint(uint64(orgId), 10)

	log.Println("[orgsUsersRepository] Updating OrgUser entry for ID (" +
		userIdStr + ") in Org ID (" + orgIdStr + ")...")

	var orgUser models.OrgUsers

	err := o.DB.Where("users_id = ? AND organization_id = ?", userId, orgId).Find(&orgUser).Error

	if err != nil {
		return orgUser, err
	}

	orgUser.Role = role

	err = o.DB.Save(&orgUser).Error

	o.DB.Preload("Users").Find(&orgUser)
	o.DB.Preload("Organization").Find(&orgUser)

	return orgUser, err
}

// Delete existing OrgUser object by ID
func (o orgUsersRepository) DeleteOrgUser(userId uint, orgId uint) error {
	userIdStr := strconv.FormatUint(uint64(userId), 10)
	orgIdStr := strconv.FormatUint(uint64(orgId), 10)

	log.Println("[orgsUsersRepository] Deleting OrgUser entry for user ID (" +
		userIdStr + ") in Org ID (" + orgIdStr + ")...")

	var orgUser models.OrgUsers
	err := o.DB.Where("users_id = ? AND organization_id = ?", userId, orgId).Find(&orgUser).Error

	if err != nil {
		return err
	}

	err = o.DB.Delete(&orgUser).Error

	return err
}
