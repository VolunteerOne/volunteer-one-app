package models

import "gorm.io/gorm"

type OrgRoles struct {
	gorm.Model
	AdminId 		uint `gorm:"not null"`
	OrganizationId 	uint `gorm:"not null"`
	Verified 		bool `gorm:"default:0;not null"`
	
	// Lower values take priority.
	// (0 = owner, 1 = manager, 2 = member) 
	Role			uint `gorm:"default:2;not null"`
}