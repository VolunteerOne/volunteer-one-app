package models

import (
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	FriendOneHandle string `gorm:"NOT NULL"`
	FriendTwoHandle string `gorm:"NOT NULL"`
	RelationshipBit string `gorm:"default:'pending'"` // other option is 'friends'
}
