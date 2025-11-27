package user

import (
	"context"
	"time"

	"github.com/Furkanberkay/todo-api-2/config"
	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo   domain.UserRepository
	config *config.Config
}

func NewUserService(repository domain.UserRepository, conf *config.Config) *Service {
	return &Service{
		repo:   repository,
		config: conf,
	}
}

func (s *Service) RegisterUser(ctx context.Context, userInput *CreateUserInput) (*domain.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.ErrInternal
	}
	hashedPasswordString := string(hashedPassword)

	user := userInput.ToModel(hashedPasswordString)

	registerErr := s.repo.RegisterUser(ctx, user)

	return user, registerErr

}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *Service) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *Service) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *Service) LoginUser(ctx context.Context, loginInput *LoginInput) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, loginInput.Email)

	if err != nil {

		return "", domain.ErrIncorrectEmailOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginInput.Password)); err != nil {
		return "", domain.ErrIncorrectEmailOrPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	secretKey := s.config.SecretKey
	if secretKey == "" {
		return "", domain.ErrInternal
	}

	tokenString, signedErr := token.SignedString([]byte(secretKey))
	if signedErr != nil {
		return "", domain.ErrInternal
	}
	return tokenString, nil

}
