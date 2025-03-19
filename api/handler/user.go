package handler

import (
	"chat_server/api/schema"
	"chat_server/middleware"
	input_port "chat_server/usecase/input_port"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUC input_port.IUserUsecase
}

func NewUserHandler(userUC input_port.IUserUsecase) *UserHandler {
	return &UserHandler{
		UserUC: userUC,
	}
}

func (h *UserHandler) FindMe(c echo.Context) error {
	user, err := middleware.GetUserFromContext(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	res := schema.User{
		UserID: user.UserID,
		UserName: user.UserName,
		Email: user.Email,
		DisplayName: user.DisplayName,
		IconUrl: user.IconUrl,
	}

	return c.JSON(http.StatusOK, res)
}