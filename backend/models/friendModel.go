package models

import (
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	id          uint `gorm:"autoincrement;primaryKey"`
	friendOneID uint
	friendTwoID uint
}
