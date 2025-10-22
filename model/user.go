package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    int            `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"not null"`
	Posts     []Post         `gorm:"foreignKey:UserID;"`
	CreatedAt time.Time      `gorm:"type:timestamptz(0)"`
	UpdatedAt time.Time      `gorm:"type:timestamptz(0)"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz(0)"`
}
