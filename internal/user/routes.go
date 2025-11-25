package user

import "github.com/labstack/echo/v4"

func (h *Handler) Routes(e *echo.Echo) {
	e.GET("register", h.RegisterUser)
}
