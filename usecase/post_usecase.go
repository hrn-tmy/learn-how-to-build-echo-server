package usecase

import "how-to-build-echo-server/repository"

type IPostUsecase interface{}

type PostUsecase struct {
	repo repository.IPostRepository
}

func NewPostUsecase(repo repository.IPostRepository) IPostUsecase {
	return PostUsecase{repo: repo}
}
