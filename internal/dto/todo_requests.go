package dto

type TodoPostRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=200"`
}

type TodoPutRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=200"`
	Completed   bool   `json:"completed" validate:"required"`
}
