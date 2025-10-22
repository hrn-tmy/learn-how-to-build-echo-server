package repository

import (
	"context"
	"how-to-build-echo-server/model"

	"gorm.io/gorm"
)

type IPostRepository interface {
	GetPosts() ([]model.Post, error)
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return PostRepository{DB: db}
}

func (r PostRepository) GetPosts() ([]model.Post, error) {
	ctx := context.Background()
	posts, err := gorm.G[model.Post](r.DB).Find(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
