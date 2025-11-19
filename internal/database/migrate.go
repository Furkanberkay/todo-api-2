package database

import (
	"log"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.Todo{}); err != nil {
		log.Fatalf("[database] automigrate failed: %v", err)
	}
	log.Printf("[database] automigrate completed")
}
