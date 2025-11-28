package todo

import (
	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"gorm.io/gorm"
)

type CreateTodoInput struct {
	Name        string
	Description string
	UserID      uint
}

type UpdateTodoInput struct {
	ID          uint
	UserID      uint
	Name        string
	Description string
	Completed   bool
}

type PatchTodoInput struct {
	ID          int
	Name        *string
	Description *string
	Completed   *bool
}

func (s *CreateTodoInput) ToModel() *domain.Todo {
	return &domain.Todo{
		Name:        s.Name,
		Description: s.Description,
		Completed:   false,
		UserID:      s.UserID,
	}
}

func (s *UpdateTodoInput) ToModel() *domain.Todo {
	return &domain.Todo{
		Model:       gorm.Model{ID: s.ID},
		Name:        s.Name,
		Description: s.Description,
		Completed:   s.Completed,
		UserID:      s.UserID,
	}
}
