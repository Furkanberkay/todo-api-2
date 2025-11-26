package user

import "github.com/labstack/echo/v4"

func (h *Handler) UserPublicRoutes(e *echo.Group) {
	e.POST("register", h.RegisterUser)
	e.POST("login", h.LoginUser)

}

func (h *Handler) UserProtectedRoutes(e *echo.Group) {
	e.POST("register", h.RegisterUser)
}
