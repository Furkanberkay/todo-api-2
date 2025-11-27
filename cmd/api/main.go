package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Furkanberkay/todo-api-2/config"
	"github.com/Furkanberkay/todo-api-2/internal/database"
	"github.com/Furkanberkay/todo-api-2/internal/todo"
	"github.com/Furkanberkay/todo-api-2/internal/user"
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.Load()

	logger := log.New(os.Stdout, "[todo] ", log.LstdFlags|log.Lshortfile)

	db := database.NewSQLite(cfg.SQLitePath)
	v := validator.New()

	todoRepo := todo.NewRepository(db, logger)
	userRepo := user.NewUserGormRepository(db, logger)

	todoService := todo.NewService(todoRepo)
	userService := user.NewUserService(userRepo, cfg)

	userHandler := user.NewHandler(userService, v)
	todoHandler := todo.NewHandler(todoService, v)

	e := echo.New()
	e.Use(echomw.Recover())
	e.Use(echomw.Logger())

	jwtConfig := echojwt.Config{
		SigningKey: []byte(cfg.SecretKey),
		ContextKey: "user",
	}

	api := e.Group("/api/v1", echojwt.WithConfig(jwtConfig))
	public := e.Group("")
	todoHandler.TodoProtectedRoutes(api)
	userHandler.UserProtectedRoutes(api)
	userHandler.UserPublicRoutes(public)

	log.Printf("[api] starting http server on %s", cfg.HTTPAddr)

	if err := e.Start(cfg.HTTPAddr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[api] server error: %v", err)
	}

}
