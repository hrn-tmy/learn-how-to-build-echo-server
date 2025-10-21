package repository

import "gorm.io/gorm"

type IPostRepository interface{}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return PostRepository{DB: db}
}
