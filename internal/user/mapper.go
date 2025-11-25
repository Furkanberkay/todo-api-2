package user

import (
	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"github.com/Furkanberkay/todo-api-2/internal/dto"
)

func MapRegisterRequestToInput(req *dto.RegisterRequest) *CreateUserInput {
	return &CreateUserInput{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}
}

func MapUserToResponse(u *domain.User) dto.RegisterResponse {
	return dto.RegisterResponse{
		ID:       u.ID,
		Name:     u.Name,
		Surname:  u.Surname,
		Email:    u.Email,
		Username: u.Username,
	}
}
