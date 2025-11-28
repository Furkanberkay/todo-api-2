package todo

import "github.com/Furkanberkay/todo-api-2/internal/domain"

type CreateTodoInput struct {
	Name        string
	Description string
}

type PatchTodoInput struct {
	ID          int
	Name        *string
	Description *string
	Completed   *bool
}

func (s *CreateTodoInput) ToModel(userID uint) *domain.Todo {
	return &domain.Todo{
		Name:        s.Name,
		Description: s.Description,
		Completed:   false,
		UserID:      userID,
	}
}
