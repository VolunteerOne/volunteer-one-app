package models

import (
	"gorm.io/gorm"
)
<<<<<<< Updated upstream

type Friend struct {
	gorm.Model
	id          uint `gorm:"autoincrement;primaryKey"`
	friendOneID uint
	friendTwoID uint
}
=======
type Friend struct {
	gorm.Model
	id       		uint   `gorm:"autoincrement;primaryKey"`
	friendOneID 	uint
	friendTwoID		uint

}
>>>>>>> Stashed changes
