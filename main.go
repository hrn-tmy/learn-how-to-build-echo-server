package main

import (
	"how-to-build-echo-server/handler"
	"how-to-build-echo-server/infra"
	"how-to-build-echo-server/repository"
	"how-to-build-echo-server/router"
	"how-to-build-echo-server/usecase"
	"log"
)

func main() {
	db, err := infra.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer infra.CloseDB(db)
	repo := repository.NewPostRepository(db)
	uc := usecase.NewPostUsecase(repo)
	h := handler.NewPostHandler(uc)
	e := router.NewRouter(h)

	e.Logger.Fatal(e.Start(":8080"))
}
