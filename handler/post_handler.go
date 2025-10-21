package handler

import (
	"how-to-build-echo-server/usecase"
	"net/http"

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
	return ctx.JSON(http.StatusOK, nil)
}

func (h PostHandler) GetPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

func (h PostHandler) CreatePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

func (h PostHandler) UpdatePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

func (h PostHandler) DeletePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
