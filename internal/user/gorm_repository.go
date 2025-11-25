package user

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	Db  *gorm.DB
	log *log.Logger
}

func NewUserGormRepository(db *gorm.DB, logger *log.Logger) domain.UserRepository {
	return &GormUserRepository{
		Db:  db,
		log: logger,
	}
}

func (r *GormUserRepository) RegisterUser(ctx context.Context, user *domain.User) error {
	result := r.Db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		errStr := result.Error.Error()

		if strings.Contains(errStr, "UNIQUE constraint failed") || strings.Contains(errStr, "Duplicate entry") {
			r.log.Printf("[user/Repo] Duplicate user: %v", result.Error)
			return domain.ErrUserAlreadyExists
		}

		r.log.Printf("[user/Repository: Create] DB Error: %v", result.Error.Error())
		return domain.ErrInternal
	}

	r.log.Printf("[user/Repository: Create] Successfully created User with ID: %d", user.ID)
	return nil
}

func (r *GormUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := new(domain.User)
	result := r.Db.WithContext(ctx).Where("email = ?", email).First(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.log.Printf("[user/Repository: GetByEmail] user Not Found for email %s", email)
			return nil, domain.ErrUserNotFound
		}
		r.log.Printf("[user/Repository: GetByEmail] General DB Error for ID %s: %v", email, result.Error.Error())
		return nil, domain.ErrInternal
	}

	return user, nil

}

func (r *GormUserRepository) DeleteUser(ctx context.Context, id uint) error {
	result := r.Db.WithContext(ctx).Where("id = ?", id).Delete(&domain.User{})

	if result.Error != nil {

		r.log.Printf("[user/Repository: DeleteUser] General DB Error for ID %d: %v", id, result.Error.Error())
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
			r.log.Printf("[user/Repository: GetByID] User Not Found for ID %d", id)
			return nil, domain.ErrUserNotFound
		}
		r.log.Printf("[user/Repository: GetByID] General DB Error for ID %d: %v", id, result.Error.Error())
		return nil, domain.ErrInternal
	}
	return user, nil
}
