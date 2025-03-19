package middleware

import (
	"chat_server/api/schema"
	"chat_server/domain/entity"
	input_port "chat_server/usecase/input_port"
	"context"
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
		token := strings.TrimPrefix(authHeader, schema.Bearer + " ")

		userID, err := a.userUC.Authenticate(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		user, err := a.userUC.FindByID(userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		setContext(c, user)

		return next(c)
	}
}

func setContext(c echo.Context, user entity.User) {
	ctx := c.Request().Context()
	ctx = setUserToContext(ctx, user)
	c.SetRequest(c.Request().WithContext(ctx))
}

type contextUserKey struct{}

func setUserToContext(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, contextUserKey{}, user)
}

func GetUserFromContext(ctx context.Context) entity.User {
	return ctx.Value(contextUserKey{}).(entity.User)
}
