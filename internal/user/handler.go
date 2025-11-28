package user

import (
	"net/http"

	"github.com/Furkanberkay/todo-api-2/internal/dto"
	"github.com/Furkanberkay/todo-api-2/internal/httpx"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service   *Service
	validator *validator.Validate
}

func NewHandler(service *Service, validate *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validate,
	}
}

func (h *Handler) RegisterUser(e echo.Context) error {
	registerDto := dto.RegisterRequest{}

	if err := e.Bind(&registerDto); err != nil {
		return httpx.InvalidBodyErr(e, err)
	}
	if err := h.validator.Struct(&registerDto); err != nil {
		validateErr := httpx.ParseValidationErrors(err)

		return e.JSON(http.StatusBadRequest, validateErr)
	}

	registerInput := MapRegisterRequestToInput(&registerDto)

	user, err := h.service.RegisterUser(e.Request().Context(), registerInput)
	if err != nil {
		return httpx.HandleServiceError(e, err)
	}

	userResponse := MapUserToResponse(user)

	return e.JSON(http.StatusCreated, userResponse)

}

func (h *Handler) LoginUser(e echo.Context) error {
	userLoginDto := new(dto.UserLoginDto)

	if err := e.Bind(userLoginDto); err != nil {
		return httpx.InvalidBodyErr(e, err)
	}

	if err := h.validator.Struct(userLoginDto); err != nil {
		parseValidateErr := httpx.ParseValidationErrors(err)
		return e.JSON(http.StatusBadRequest, parseValidateErr)
	}

	loginInput := LoginInput{
		Email:    userLoginDto.Email,
		Password: userLoginDto.Password,
	}

	token, err := h.service.LoginUser(e.Request().Context(), &loginInput)

	if err != nil {
		return httpx.HandleServiceError(e, err)
	}

	resp := dto.LoginResponse{
		Token: token,
	}

	return e.JSON(http.StatusOK, resp)
}
