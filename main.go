package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type post struct {
  postID int `gorm:"primaryKey"`
  title string
  body string
  createdAt time.Time
  updateAt time.Time
  deletedAt gorm.DeletedAt
}

type dbConfig struct {
  user string
  password string
  host string
  port string
  name string
}

func newDB() (*gorm.DB, error) {
  if err := godotenv.Load(); err != nil {
    log.Fatal(err)
  }
  cfg := dbConfig{
    user: os.Getenv("DB_USER"),
    password: os.Getenv("DB_PASSWORD"),
    host: os.Getenv("DB_HOST"),
    port: os.Getenv("DB_PORT"),
    name: os.Getenv("DB_NAME"),
  }
  dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.user, cfg.password, cfg.host, cfg.port, cfg.name)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }

  return db, err
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Go!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}