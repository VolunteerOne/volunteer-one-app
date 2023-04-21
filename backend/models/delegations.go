package models

import (
	"gorm.io/gorm"
)

type Delegations struct {
	gorm.Model
	RefreshToken string
	UsersID      uint `gorm:"unique;not null"`

	Users Users `gorm:"foreignkey:UsersID"`
}
