package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	PostID    int `gorm:"primaryKey"`
	UserID    int
	Title     string         `gorm:"not null"`
	Body      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"type:datetime"`
	UpdatedAt time.Time      `gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime"`
}
