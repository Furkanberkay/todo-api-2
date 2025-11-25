package dto

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Surname  string `json:"surname" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}
