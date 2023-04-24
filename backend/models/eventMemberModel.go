package models

import "gorm.io/gorm"

type EventMember struct {
	gorm.Model
	UserId uint 
	EventId uint

	User Users `gorm:"foreignkey:UserId"`
	Event Event `gorm:"foreignkey:EventId"`
}