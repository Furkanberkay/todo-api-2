package todo

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"github.com/Furkanberkay/todo-api-2/internal/dto"
	"github.com/Furkanberkay/todo-api-2/internal/httpx"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ResponseErr struct {
	Message string `json:"message"`
}

type Handler struct {
	service   *Service
	validator *validator.Validate
}

func NewHandler(service *Service, validator *validator.Validate) *Handler {
	return &Handler{service: service, validator: validator}
}

func (h *Handler) GetTodos(e echo.Context) error {

	var todoList []dto.TodoListItemResponse

	todos, err := h.service.GetTodos(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ResponseErr{Message: domain.ErrInternal.Error()})
	}

	for _, todo := range todos {
		todoList = append(todoList, dto.TodoListItemResponse{
			ID:          todo.ID,
			Name:        todo.Name,
			Description: todo.Description,
			Completed:   todo.Completed,
		})
	}

	return e.JSON(http.StatusOK, todoList)
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

	todoResp := dto.TodoDetailResponse{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
	return e.JSON(http.StatusOK, todoResp)
}

func (h *Handler) CreateTodo(e echo.Context) error {
	createTodoDTO := dto.TodoPostRequest{}

	if err := e.Bind(&createTodoDTO); err != nil {
		return e.JSON(http.StatusBadRequest, ResponseErr{
			Message: "invalid request body",
		})
	}

	if err := h.validator.Struct(&createTodoDTO); err != nil {
		validationErrorResponse := httpx.ParseValidationErrors(err)
		return e.JSON(http.StatusBadRequest, validationErrorResponse)
	}
	todo := CreateTodoInput{
		Name:        createTodoDTO.Name,
		Description: createTodoDTO.Description,
	}

	createdTodo, err := h.service.CreateTodo(e.Request().Context(), &todo)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ResponseErr{
			Message: domain.ErrInternal.Error(),
		})
	}
	todoResp := dto.TodoDetailResponse{
		ID:          createdTodo.ID,
		Name:        createdTodo.Name,
		Description: createdTodo.Description,
		Completed:   createdTodo.Completed,
		CreatedAt:   createdTodo.CreatedAt,
	}

	return e.JSON(http.StatusCreated, todoResp)

}

func (h *Handler) DeleteTodo(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, ResponseErr{
			Message: "id must be a number",
		})
	}
	if err := h.service.DeleteTodo(e.Request().Context(), id); err != nil {
		if errors.Is(err, domain.ErrTodoNotFound) {
			return e.JSON(http.StatusNotFound, ResponseErr{
				Message: domain.ErrTodoNotFound.Error(),
			})
		}
		return e.JSON(http.StatusInternalServerError, ResponseErr{
			Message: domain.ErrInternal.Error(),
		})
	}
	return e.NoContent(http.StatusNoContent)
}

func (h *Handler) UpdateTodo(e echo.Context) error {
	todoPutRequest := dto.TodoPutRequest{}

	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, ResponseErr{
			Message: "id must be a number",
		})
	}

	if err := e.Bind(&todoPutRequest); err != nil {
		return e.JSON(http.StatusBadRequest, ResponseErr{
			Message: "invalid request body",
		})
	}

	if err := h.validator.Struct(&todoPutRequest); err != nil {
		validationErrResponse := httpx.ParseValidationErrors(err)
		return e.JSON(http.StatusBadRequest, validationErrResponse)
	}

	todo := domain.Todo{}

	todo.ID = uint(id)
	todo.Name = todoPutRequest.Name
	todo.Description = todoPutRequest.Description
	todo.Completed = todoPutRequest.Completed

	if err := h.service.UpdateTodo(e.Request().Context(), &todo); err != nil {
		if errors.Is(err, domain.ErrTodoNotFound) {
			return e.JSON(http.StatusNotFound, ResponseErr{
				Message: domain.ErrTodoNotFound.Error(),
			})
		}
		return e.JSON(http.StatusInternalServerError, ResponseErr{
			Message: domain.ErrInternal.Error(),
		})
	}

	todoResponse := dto.TodoDetailResponse{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Completed:   todo.Completed,
		UpdatedAt:   todo.UpdatedAt,
	}

	return e.JSON(http.StatusOK, todoResponse)

}
