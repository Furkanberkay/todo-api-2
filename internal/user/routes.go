package user

import "github.com/labstack/echo/v4"

func (h *Handler) UserRoutes(e *echo.Echo) {
	e.POST("register", h.RegisterUser)
}
