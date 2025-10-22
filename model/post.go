package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	PostID    int `gorm:"primaryKey"`
	Title     string         `gorm:"not null"`
	Body      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"type:timestamptz(0)"`
	UpdatedAt time.Time      `gorm:"type:timestamptz(0)"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz(0)"`
}
