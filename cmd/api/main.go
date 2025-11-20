package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Furkanberkay/todo-api-2/config"
	"github.com/Furkanberkay/todo-api-2/internal/database"
	"github.com/Furkanberkay/todo-api-2/internal/todo"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Load()

	logger := log.New(os.Stdout, "[todo] ", log.LstdFlags|log.Lshortfile)

	db := database.NewSQLite(cfg.SQLitePath)

	repo := todo.NewRepository(db, logger)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	e := echo.New()

	handler.RegisterRoutes(e)

	log.Printf("[api] starting http server on %s", cfg.HTTPAddr)

	if err := e.Start(cfg.HTTPAddr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[api] server error: %v", err)
	}
}
