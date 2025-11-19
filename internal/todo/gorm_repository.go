package todo

import (
	"context"
	"errors"
	"log"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	Db  *gorm.DB
	log *log.Logger
}

func NewRepository(db *gorm.DB, log *log.Logger) domain.TodoRepository {
	return &Repository{Db: db, log: log}
}

func (r *Repository) GetTodos(ctx context.Context) ([]domain.Todo, error) {
	var todos []domain.Todo
	result := r.Db.WithContext(ctx).Find(&todos)

	if result.Error != nil {
		r.log.Printf("[todo/Repository: GetTodos] DB Error: %v", result.Error.Error())
		return nil, domain.ErrInternal
	}

	return todos, nil
}

func (r *Repository) GetTodoByID(ctx context.Context, id int) (*domain.Todo, error) {
	todo := new(domain.Todo)
	result := r.Db.WithContext(ctx).First(todo, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.log.Printf("[todo/Repository: GetByID] Todo Not Found for ID %d", id)
			return nil, domain.ErrTodoNotFound
		}
		r.log.Printf("[todo/Repository: GetByID] General DB Error for ID %d: %v", id, result.Error.Error())
		return nil, domain.ErrInternal
	}
	return todo, nil
}

func (r *Repository) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	result := r.Db.WithContext(ctx).Create(todo)

	if result.Error != nil {
		r.log.Printf("[todo/Repository: Create] DB Error: %v", result.Error.Error())
		return domain.ErrInternal
	}

	r.log.Printf("[todo/Repository: Create] Successfully created Todo with ID: %d", todo.ID)
	return nil
}

func (r *Repository) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	result := r.Db.WithContext(ctx).Where("id = ?", todo.ID).Updates(todo)

	if result.Error != nil {
		r.log.Printf("[todo/Repository: UpdateTodo] General DB Error for ID %d: %v", todo.ID, result.Error.Error())
		return domain.ErrInternal
	}

	if result.RowsAffected == 0 {
		r.log.Printf("[todo/Repository: UpdateTodo] Todo Not Found for ID %d", todo.ID)
		return domain.ErrTodoNotFound
	}

	r.log.Printf("[todo/Repository: UpdateTodo] Successfully updated Todo with ID: %d", todo.ID)
	return nil
}

func (r *Repository) DeleteTodo(ctx context.Context, id int) error {
	result := r.Db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Todo{})

	if result.Error != nil {

		r.log.Printf("[todo/Repository: DeleteTodo] General DB Error for ID %d: %v", id, result.Error.Error())
		return domain.ErrInternal
	}

	if result.RowsAffected == 0 {
		return domain.ErrTodoNotFound
	}
	return nil
}
