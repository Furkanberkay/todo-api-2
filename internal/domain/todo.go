package domain

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string `gorm:"not null,size:200"`
	Description string `gorm:"not null,size:500"`
	Completed   bool   `gorm:"default:false"`
}

func (Todo) TableName() string {
	return "todos"
}

type TodoRepository interface {
	GetTodos(ctx context.Context, page int, limit int) ([]Todo, int, error)
	GetTodoByID(ctx context.Context, id int) (*Todo, error)
	CreateTodo(ctx context.Context, todo *Todo) error
	UpdateTodo(ctx context.Context, todo *Todo) error
	DeleteTodo(ctx context.Context, id int) error
}

var ErrTodoNotFound = errors.New("todo not found")
var ErrInternal = errors.New("server internal error")
var ErrValidation = errors.New("validation Error")
