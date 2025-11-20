package todo

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/todos", h.GetTodos)
	e.GET("/todos/:id", h.GetTodoByID)
	e.POST("/todos", h.CreateTodo)
	e.PUT("/todos/:id", h.UpdateTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)

}
