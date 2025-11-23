package dto

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Surname  string `json:"surname" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
