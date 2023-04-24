package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Organization    Organization `gorm:"foreignkey:OrganizationID"`
	OrganizationID  uint
	Name        	string
	Address			string
	Date 			time.Time
	Description 	string
	Interests		string
	Skills			string
	GoodFor			string
	CauseAreas		string
	Requirements 	string
}
