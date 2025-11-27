package todo

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	Db     *gorm.DB
	logger *slog.Logger
}

func NewRepository(db *gorm.DB, logger *slog.Logger) domain.TodoRepository {
	return &Repository{Db: db, logger: logger}
}

func (r *Repository) GetTodos(ctx context.Context, page int, limit int) ([]domain.Todo, int, error) {
	var todos []domain.Todo
	var totalCount int64

	offset := (page - 1) * limit

	if err := r.Db.WithContext(ctx).Model(&domain.Todo{}).Count(&totalCount).Error; err != nil {
		r.logger.Error("database error during todo count",
			slog.String("component", "TodoRepository"),
			slog.String("error", err.Error()),
		)
		return nil, 0, domain.ErrInternal
	}

	result := r.Db.WithContext(ctx).Offset(offset).Limit(limit).Find(&todos)

	if result.Error != nil {
		r.logger.Error("database error during todo list fetch",
			slog.String("component", "TodoRepository"),
			slog.String("error", result.Error.Error()),
		)
		return nil, 0, domain.ErrInternal
	}

	return todos, int(totalCount), nil
}

func (r *Repository) GetTodoByID(ctx context.Context, id int) (*domain.Todo, error) {
	todo := new(domain.Todo)
	result := r.Db.WithContext(ctx).First(todo, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrTodoNotFound
		}

		r.logger.Error("database error during todo lookup by id",
			slog.String("component", "TodoRepository"),
			slog.Int("todo_id", id),
			slog.String("error", result.Error.Error()),
		)
		return nil, domain.ErrInternal
	}
	return todo, nil
}

func (r *Repository) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	result := r.Db.WithContext(ctx).Create(todo)

	if result.Error != nil {
		r.logger.Error("database error during todo creation",
			slog.String("component", "TodoRepository"),
			slog.String("error", result.Error.Error()),
		)
		return domain.ErrInternal
	}

	return nil
}

func (r *Repository) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	result := r.Db.WithContext(ctx).Model(&domain.Todo{}).Where("id = ?", todo.ID).Updates(todo)

	if result.Error != nil {
		r.logger.Error("database error during todo update",
			slog.String("component", "TodoRepository"),
			slog.Int("todo_id", int(todo.ID)),
			slog.String("error", result.Error.Error()),
		)
		return domain.ErrInternal
	}

	if result.RowsAffected == 0 {
		return domain.ErrTodoNotFound
	}

	return nil
}

func (r *Repository) DeleteTodo(ctx context.Context, id int) error {
	result := r.Db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Todo{})

	if result.Error != nil {
		r.logger.Error("database error during todo deletion",
			slog.String("component", "TodoRepository"),
			slog.Int("todo_id", id),
			slog.String("error", result.Error.Error()),
		)
		return domain.ErrInternal
	}

	if result.RowsAffected == 0 {
		return domain.ErrTodoNotFound
	}
	return nil
}
