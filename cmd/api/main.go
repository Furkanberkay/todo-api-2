package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Furkanberkay/todo-api-2/config"
	"github.com/Furkanberkay/todo-api-2/internal/database"
	"github.com/Furkanberkay/todo-api-2/internal/middleware"
	"github.com/Furkanberkay/todo-api-2/internal/todo"
	"github.com/Furkanberkay/todo-api-2/internal/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.Load()

	handler := slog.NewJSONHandler(os.Stdout, nil)
	slogLogger := slog.New(handler)

	db := database.NewSQLite(cfg.SQLitePath)
	v := validator.New()

	todoRepo := todo.NewRepository(db, slogLogger)
	userRepo := user.NewUserGormRepository(db, slogLogger)

	todoService := todo.NewService(todoRepo)
	userService := user.NewUserService(userRepo, cfg)

	userHandler := user.NewHandler(userService, v)
	todoHandler := todo.NewHandler(todoService, v, slogLogger)

	e := echo.New()
	e.Use(echomw.Recover())
	e.Use(echomw.Logger())

	authMiddleware := middleware.NewAuthMiddleware(cfg.SecretKey, slogLogger)

	api := e.Group("/api/v1", authMiddleware)
	public := e.Group("")
	todoHandler.TodoProtectedRoutes(api)
	userHandler.UserProtectedRoutes(api)
	userHandler.UserPublicRoutes(public)

	log.Printf("[api] starting http server on %s", cfg.HTTPAddr)

	if err := e.Start(cfg.HTTPAddr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[api] server error: %v", err)
	}

}
