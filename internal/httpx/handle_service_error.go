package httpx

import (
	"errors"
	"net/http"

	"github.com/Furkanberkay/todo-api-2/internal/domain"
	"github.com/labstack/echo/v4"
)

type ResponseErr struct {
	Message string `json:"message"`
}

func HandleServiceError(c echo.Context, err error) error {
	if errors.Is(err, domain.ErrTodoNotFound) {
		return c.JSON(http.StatusNotFound, ResponseErr{Message: domain.ErrTodoNotFound.Error()})
	}

	return c.JSON(http.StatusInternalServerError, ResponseErr{Message: domain.ErrInternal.Error()})
}

func IdMapError(c echo.Context, err error) error {

	return c.JSON(http.StatusBadRequest, ResponseErr{Message: "id must be a number"})
}

func InvalidBodyErr(c echo.Context, err error) error {

	return c.JSON(http.StatusBadRequest, ResponseErr{Message: "invalid request body"})
}
