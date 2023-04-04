package models

<<<<<<< HEAD
import (
	"gorm.io/gorm"
)
=======
import "gorm.io/gorm"
>>>>>>> 4e9ea3a (Update (#5))

type Object struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Value string
}
