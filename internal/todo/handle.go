package todo

import (
	"net/http"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"github.com/labstack/echo/v4"
)

type ResponseErr struct {
	Message string `json:"message"`
}

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTodos(e echo.Context) error {

	todos, err := h.service.GetTodos(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ResponseErr{Message: domain.ErrInternal.Error()})
	}
	return e.JSON(http.StatusOK, todos)
}
