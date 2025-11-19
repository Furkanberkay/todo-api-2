package todo

import (
	"context"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
)

type Service struct {
	repo domain.TodoRepository
}

func NewService(repo domain.TodoRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTodos(ctx context.Context) ([]domain.Todo, error) {
	return s.repo.GetTodos(ctx)
}

func (s *Service) GetTodoByID(ctx context.Context, id int) (*domain.Todo, error) {
	return s.repo.GetTodoByID(ctx, id)
}

func (s *Service) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	return s.repo.CreateTodo(ctx, todo)
}
func (s *Service) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	return s.repo.UpdateTodo(ctx, todo)
}
func (s *Service) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.DeleteTodo(ctx, id)
}
