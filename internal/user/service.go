package user

import (
	"context"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) RegisterUser(ctx context.Context, userInput *CreateUserInput) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.ErrInternal
	}

	user := domain.User{
		Name:         userInput.Name,
		Email:        userInput.Email,
		Surname:      userInput.Surname,
		Username:     userInput.Username,
		PasswordHash: string(hashedPassword),
	}

	return s.repo.RegisterUser(ctx, &user)

}
