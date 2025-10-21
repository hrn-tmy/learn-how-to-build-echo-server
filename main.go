package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 必要なこと
// DB接続
// Repository層の作成
// Usecase層の作成
// Handler層の作成
// Router層の作成
// 今回はTODOリストを作成する前提でいいかも
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Go!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}