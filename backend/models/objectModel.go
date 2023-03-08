package models

import (
	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Value string
}

// type Users struct {
// 	gorm.Model
// 	id       uint   `gorm:"unique; autoincrement;primaryKey"`
// 	handle   string `gorm: "not null"`
// 	email    string `gorm: "NOT NULL"`
// 	password string `gorm: "NOT NULL"`
// 	// birthdate datatypes.Date `gorm: "NOT NULL"`
// 	birthdate string `gorm: "NOT NULL"`
// 	firstName string `gorm: "NOT NULL"`
// 	lastName  string `gorm: "NOT NULL"`
// 	// profilePic mediumblob,
// 	interests string
// 	verified  uint
// }
