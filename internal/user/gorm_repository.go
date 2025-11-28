package user

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	Db     *gorm.DB
	logger *slog.Logger
}

func NewUserGormRepository(db *gorm.DB, logger *slog.Logger) domain.UserRepository {
	return &GormUserRepository{
		Db:     db,
		logger: logger,
	}
}

func (r *GormUserRepository) RegisterUser(ctx context.Context, user *domain.User) error {
	result := r.Db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		errStr := result.Error.Error()

		if strings.Contains(errStr, "UNIQUE constraint failed") || strings.Contains(errStr, "Duplicate entry") {
			return domain.ErrUserAlreadyExists
		}

		r.logger.Error("database error during user creation",
			slog.String("component", "UserRepository"),
			slog.String("error", errStr),
		)
		return domain.ErrInternal
	}

	return nil
}

func (r *GormUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := new(domain.User)
	result := r.Db.WithContext(ctx).Where("email = ?", email).First(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		r.logger.Error("database error during user lookup by email",
			slog.String("component", "UserRepository"),
			slog.String("error", result.Error.Error()),
		)
		return nil, domain.ErrInternal
	}

	return user, nil

}

func (r *GormUserRepository) DeleteUser(ctx context.Context, id uint) error {
	result := r.Db.WithContext(ctx).Where("id = ?", id).Delete(&domain.User{})

	if result.Error != nil {

		r.logger.Error("database error during user deletion",
			slog.String("component", "UserRepository"),
			slog.Int("user_id", int(id)),
			slog.String("error", result.Error.Error()),
		)
		return domain.ErrInternal
	}

	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *GormUserRepository) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	user := new(domain.User)

	result := r.Db.WithContext(ctx).First(user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		r.logger.Error("database error during user lookup by id",
			slog.String("component", "UserRepository"),
			slog.Int("user_id", int(id)),
			slog.String("error", result.Error.Error()),
		)
		return nil, domain.ErrInternal
	}
	return user, nil
}
