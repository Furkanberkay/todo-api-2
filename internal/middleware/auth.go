package middleware

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const contextKey = "user"
const ClaimUserIDKey = "user_id"

func NewAuthMiddleware(secretKey string, logger *slog.Logger) echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(secretKey),
		ContextKey: contextKey,
		ErrorHandler: func(c echo.Context, err error) error {

			logger.Warn("authentication failed",
				slog.String("error", err.Error()),
				slog.String("ip", c.RealIP()),
			)
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "unauthorized access",
			})
		},
	}
	return echojwt.WithConfig(config)
}

func GetUserID(e echo.Context) (uint, error) {
	tokenData := e.Get(contextKey)
	if tokenData == nil {
		return 0, errors.New("token not found in context")
	}

	t, ok := tokenData.(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token format in context")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userIdFloat, ok := claims[ClaimUserIDKey].(float64)
	if !ok {
		return 0, fmt.Errorf("claim '%s' missing or invalid", ClaimUserIDKey)
	}

	id := uint(userIdFloat)

	return id, nil

}
