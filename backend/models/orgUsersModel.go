package models

import "gorm.io/gorm"

type OrgUsers struct {
	gorm.Model
	UsersID        uint `gorm:"not null" binding:"required"`
	OrganizationID uint `gorm:"not null" binding:"required"`
	Verified       bool `gorm:"default:0;not null"`

	// Lower values take priority.
	// We have leeway for additional roles.
	// (0 = owner, 1 = manager, 10 = member)
	Role uint `gorm:"default:10;not null"`

	Users        Users        `gorm:"foreignkey:UsersID"`
	Organization Organization `gorm:"foreignkey:OrganizationID"`
}
