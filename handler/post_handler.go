package handler

import (
	"how-to-build-echo-server/model"
	"how-to-build-echo-server/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IPostHandler interface {
	GetPosts(ctx echo.Context) error
	GetPost(ctx echo.Context) error
	CreatePost(ctx echo.Context) error
	UpdatePost(ctx echo.Context) error
	DeletePost(ctx echo.Context) error
}

type PostHandler struct {
	uc usecase.IPostUsecase
}

func NewPostHandler(uc usecase.IPostUsecase) IPostHandler {
	return PostHandler{uc: uc}
}

func (h PostHandler) GetPosts(ctx echo.Context) error {
	posts, err := h.uc.GetPosts()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, posts)
}

func (h PostHandler) GetPost(ctx echo.Context) error {
  strPostID := ctx.Param("id")
  postID, err := strconv.Atoi(strPostID)
  if err != nil {
    return ctx.JSON(http.StatusInternalServerError, err)
  }
  post, err := h.uc.GetPost(postID)
  if post.PostID == 0 {
    return ctx.JSON(http.StatusNotFound, err)
  }
  if err != nil {
    return ctx.JSON(http.StatusInternalServerError, err)
  }
	return ctx.JSON(http.StatusOK, post)
}

func (h PostHandler) CreatePost(ctx echo.Context) error {
  var data model.Post
  if err := ctx.Bind(&data); err != nil {
    return ctx.JSON(http.StatusBadRequest, err)
  }
  if err := h.uc.CreatePost(data); err != nil {
    return ctx.JSON(http.StatusInternalServerError, err)
  }
	return ctx.JSON(http.StatusCreated, nil)
}

func (h PostHandler) UpdatePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

func (h PostHandler) DeletePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
