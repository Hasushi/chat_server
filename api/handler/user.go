package handler

import (
	"chat_server/api/schema"
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

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req schema.CreateUserReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}
	
	user, err := h.UserUC.Create(req.UserName, req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(500, "Internal server error")
	}

	res := schema.CreateUserRes{
		UserID:       user.UserID,
		UserName: user.UserName,
		Email:    user.Email,
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req schema.LoginReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}

	token, expiresAt, err := h.UserUC.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(500, "Internal server error")
	}

	res := schema.LoginRes{
		Token: token,
		ExpiresAt: int64(expiresAt),
	}

	return c.JSON(http.StatusOK, res)
}