package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Handle          string `gorm:"NOT NULL"`
	PostDescription string
	Likes           uint `gorm:"default:0"`
}

type Comments struct {
	gorm.Model
	PostsID            uint   `gorm:"NOT NULL"`
	Handle             string `gorm:"NOT NULL"`
	CommentDescription string `gorm:"NOT NULL"`

	Posts Posts `gorm:"foreignkey:PostsID"`
}

type Likes struct {
	gorm.Model
	PostsID uint   `gorm:"NOT NULL"`
	Handle  string `gorm:"NOT NULL"`

	Posts Posts `gorm:"foreignkey:PostsID"`
}
