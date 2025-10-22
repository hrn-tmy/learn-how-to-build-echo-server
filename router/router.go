package router

import (
	"how-to-build-echo-server/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(h handler.IPostHandler) *echo.Echo {
	e := echo.New()

	e.GET("/posts", h.GetPosts)
	e.GET("/post/:id", h.GetPost)
	e.POST("/post", h.CreatePost)
	e.PUT("/post/:id", h.UpdatePost)
	e.DELETE("/post/:id", h.DeletePost)

	return e
}
