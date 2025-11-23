package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Furkanberkay/todo-api-2/config"
	"github.com/Furkanberkay/todo-api-2/internal/database"
	"github.com/Furkanberkay/todo-api-2/internal/todo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	cfg := config.Load()

	mypassword := "berkay2001"
	hashed, err := bcrypt.GenerateFromPassword([]byte(mypassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error")
		return
	}
	hashedString := string(hashed)
	fmt.Println(hashedString)

	logger := log.New(os.Stdout, "[todo] ", log.LstdFlags|log.Lshortfile)

	db := database.NewSQLite(cfg.SQLitePath)
	v := validator.New()

	repo := todo.NewRepository(db, logger)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service, v)

	e := echo.New()
	e.Use(middleware.Recover())

	handler.RegisterRoutes(e)

	log.Printf("[api] starting http server on %s", cfg.HTTPAddr)

	if err := e.Start(cfg.HTTPAddr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[api] server error: %v", err)
	}

}
