package usecase

import (
	"how-to-build-echo-server/model"
	"how-to-build-echo-server/repository"
)

type IPostUsecase interface {
	GetPosts() ([]model.Post, error)
  GetPost(postID int) (model.Post, error)
  CreatePost(data model.Post) error
  UpdatePost(data model.Post, postID int) (int, error)
  DeletePost(postID int) (int, error)
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

func (uc PostUsecase) GetPost(postID int) (model.Post, error) {
  post, err := uc.repo.GetPost(postID)
  if err != nil {
    return model.Post{}, err
  }
  return post, nil
}

func (uc PostUsecase) CreatePost(data model.Post) error {
  if err := uc.repo.CreatePost(data); err != nil {
    return err
  }
  return nil
}

func (uc PostUsecase) UpdatePost(data model.Post, postID int) (int, error) {
  row, err := uc.repo.UpdatePost(data, postID)
  if err != nil {
    return row, err
  }
  return row, nil
}

func (uc PostUsecase) DeletePost(postID int) (int, error) {
  row, err := uc.repo.DeletePost(postID)
  if err != nil {
    return row, err
  }
  return row, nil
}