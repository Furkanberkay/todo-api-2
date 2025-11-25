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
		c.Logger().Warnf("Service Warning (Not Found): %v", err)
		return c.JSON(http.StatusNotFound, ResponseErr{Message: domain.ErrTodoNotFound.Error()})
	}

	if errors.Is(err, domain.ErrUserAlreadyExists) {
		return c.JSON(http.StatusConflict, ResponseErr{Message: domain.ErrUserAlreadyExists.Error()})
	}

	c.Logger().Errorf("Service Internal Error: %v", err)
	return c.JSON(http.StatusInternalServerError, ResponseErr{Message: domain.ErrInternal.Error()})
}

func IdMapError(c echo.Context, err error) error {
	c.Logger().Errorf("ID Mapping Error: %v", err)
	return c.JSON(http.StatusBadRequest, ResponseErr{Message: "id must be a number"})
}

func InvalidBodyErr(c echo.Context, err error) error {
	c.Logger().Errorf("Body Parse Error: %v", err)
	return c.JSON(http.StatusBadRequest, ResponseErr{Message: "invalid request body"})
}
