package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Handle          string `gorm:"NOT NULL"`
	PostDescription string
	Likes           uint `gorm:"default:0"`
	CommentId       uint
}

type Comments struct {
	PostID             uint
	Handle             string
	CommentDescription string
}

type Likes struct {
	PostID uint
	Handle string
}
