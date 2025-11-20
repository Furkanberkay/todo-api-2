package todo

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *Handler) GetTodoByID(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, ResponseErr{Message: "id must be a number"})
	}
	todo, errService := h.service.GetTodoByID(e.Request().Context(), id)

	if errService != nil {
		if errors.Is(errService, domain.ErrTodoNotFound) {
			return e.JSON(http.StatusNotFound, ResponseErr{
				Message: domain.ErrTodoNotFound.Error(),
			})
		}
		return e.JSON(http.StatusInternalServerError, ResponseErr{
			Message: domain.ErrInternal.Error(),
		})
	}
	return e.JSON(http.StatusOK, todo)
}
