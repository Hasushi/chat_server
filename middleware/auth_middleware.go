package middleware

import (
	input_port "chat_server/usecase/input_port"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	userUC input_port.IUserUsecase
}

func NewAuthMiddleware(userUC input_port.IUserUsecase) AuthMiddleware {
	return AuthMiddleware{
		userUC: userUC,
	}
}

func (a *AuthMiddleware) Authenticate(next echo.HandlerFunc) (echo.HandlerFunc){
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := a.userUC.Authenticate(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		_, err = a.userUC.FindByID(userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}