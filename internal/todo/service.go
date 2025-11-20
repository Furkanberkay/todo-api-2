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

func (s *Service) CreateTodo(ctx context.Context, todoInput *CreateTodoInput) (*domain.Todo, error) {

	todo := &domain.Todo{
		Name:        todoInput.Name,
		Description: todoInput.Description,
		Completed:   false,
	}

	if err := s.repo.CreateTodo(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Service) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	return s.repo.UpdateTodo(ctx, todo)
}

func (s *Service) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.DeleteTodo(ctx, id)
}

func (s *Service) PatchTodo(ctx context.Context, patchTodo PatchTodoInput) (*domain.Todo, error) {

	if patchTodo.Name == nil && patchTodo.Description == nil && patchTodo.Completed == nil {
		return nil, domain.ErrValidation
	}
	todo, err := s.GetTodoByID(ctx, patchTodo.ID)
	if err != nil {
		return nil, err
	}
	if patchTodo.Name != nil {
		todo.Name = *patchTodo.Name
	}
	if patchTodo.Description != nil {
		todo.Description = *patchTodo.Description
	}
	if patchTodo.Completed != nil {
		todo.Completed = *patchTodo.Completed
	}

	if err := s.repo.UpdateTodo(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil

}
