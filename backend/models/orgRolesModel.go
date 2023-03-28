package models

import "gorm.io/gorm"

type OrgRoles struct {
	gorm.Model
	UsersID        uint `gorm:"not null"`
	OrganizationID uint `gorm:"not null"`
	Verified       bool `gorm:"default:0;not null"`

	// Lower values take priority.
	// We have leeway for additional roles.
	// (0 = owner, 1 = manager, 10 = member)
	Role uint `gorm:"default:5;not null"`

	Users        Users        `gorm:"foreignkey:UsersID"`
	Organization Organization `gorm:"foreignkey:OrganizationID"`
}
