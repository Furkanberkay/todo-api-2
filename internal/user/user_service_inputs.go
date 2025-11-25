package user

import "github.com/Furkanberkay/todo-api-2/internal/domain"

type CreateUserInput struct {
	Name     string
	Surname  string
	Email    string
	Username string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

func (s *CreateUserInput) ToModel(hashedPassword string) *domain.User {
	return &domain.User{
		Name:         s.Name,
		Username:     s.Username,
		PasswordHash: hashedPassword,
		Surname:      s.Surname,
		Email:        s.Email,
	}
}
