package models

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Value string
}
