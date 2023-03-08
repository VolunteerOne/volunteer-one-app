package models

import (
	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Value string
}

type Users struct {
	gorm.Model
	id       uint   `gorm:"unique; autoincrement;primaryKey"`
	Handle   string `gorm: "not null"`
	Email    string `gorm: "NOT NULL"`
	Password string `gorm: "NOT NULL"`
	// birthdate datatypes.Date `gorm: "NOT NULL"`
	Birthdate string `gorm: "NOT NULL"`
	FirstName string `gorm: "NOT NULL"`
	LastName  string `gorm: "NOT NULL"`
	// profilePic mediumblob,
	Interests string
	Verified  uint
}
