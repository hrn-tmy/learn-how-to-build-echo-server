package handler

import "how-to-build-echo-server/usecase"

type IPostHandler interface{}

type PostHandler struct {
	uc usecase.IPostUsecase
}

func NewPostHandler(uc usecase.IPostUsecase) IPostHandler {
	return PostHandler{uc: uc}
}