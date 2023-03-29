package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Handle   string `gorm:"unique,not null"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	// birthdate datatypes.Date `gorm: "NOT NULL"`
	Birthdate string `gorm:"NOT NULL" json:"bday"`
	FirstName string `gorm:"NOT NULL" json:"first"`
	LastName  string `gorm:"NOT NULL" json:"last"`
	// profilePic mediumblob,
	Interests string
	Verified  uint
	// Password forgotten reset code
	ResetCode uuid.UUID
}
