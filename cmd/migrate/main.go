package main

import (
	"log"

	"github.com/Furkanberkay/todo-api-2/config"
	"github.com/Furkanberkay/todo-api-2/internal/database"
)

func main() {
	cfg := config.Load()

	db := database.NewSQLite(cfg.SQLitePath)

	database.AutoMigrate(db)

	log.Println("[migrate] done ✅")
}
