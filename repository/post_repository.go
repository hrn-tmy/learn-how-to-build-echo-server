package repository

import (
	"context"
	"errors"
	"fmt"
	"how-to-build-echo-server/model"

	"gorm.io/gorm"
)

type IPostRepository interface {
	GetPosts() ([]model.Post, error)
  GetPost(postID int) (model.Post, error)
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

func (r PostRepository) GetPost(postID int) (model.Post, error) {
  ctx := context.Background()
  post, err := gorm.G[model.Post](r.DB).Where("post_id = ?", postID).Take(ctx)
  if err != nil {
    return model.Post{}, err
  }
  if errors.Is(err, gorm.ErrRecordNotFound) {
    return model.Post{}, fmt.Errorf("レコードが存在しません。")
  }
  return post, nil
}