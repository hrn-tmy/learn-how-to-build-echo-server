package usecase

import (
	"how-to-build-echo-server/model"
	"how-to-build-echo-server/repository"
)

type IPostUsecase interface {
	GetPosts() ([]model.Post, error)
}

type PostUsecase struct {
	repo repository.IPostRepository
}

func NewPostUsecase(repo repository.IPostRepository) IPostUsecase {
	return PostUsecase{repo: repo}
}

func (uc PostUsecase) GetPosts() ([]model.Post, error) {
	posts, err := uc.repo.GetPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
