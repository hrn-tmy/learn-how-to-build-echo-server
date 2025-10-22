package main

import (
	"context"
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

type Post struct {
  PostID int `gorm:"primaryKey"`
  Title string
  Body string
  CreatedAt time.Time
  UpdateAt time.Time
  DeletedAt gorm.DeletedAt
}

type DBConfig struct {
  User string
  Password string
  Host string
  Port string
  Name string
}

func newDB() (*gorm.DB, error) {
  if err := godotenv.Load(); err != nil {
    log.Fatal(err)
  }
  cfg := DBConfig{
    User: os.Getenv("DB_USER"),
    Password: os.Getenv("DB_PASSWORD"),
    Host: os.Getenv("DB_HOST"),
    Port: os.Getenv("DB_PORT"),
    Name: os.Getenv("DB_NAME"),
  }
  dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }

  return db, err
}

func main() {
	e := echo.New()
  db, err := newDB()
  if err != nil {
    log.Fatal(err)
  }
  sqlDB, err := db.DB()
  if err != nil {
    log.Fatal(err)
  }
  if err := sqlDB.Close(); err != nil {
    log.Fatal(err)
  }

	// 一覧取得
  e.GET("/posts", func(ctx echo.Context) error {
    c := context.Background()
    posts, err := gorm.G[Post](db).Find(c)
    if err != nil {
      return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
		return ctx.JSON(http.StatusOK, posts)
	})

	e.Logger.Fatal(e.Start(":8080"))
}