package todo

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Routes(e *echo.Echo) {
	e.GET("/todos", h.GetTodos)
}
