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

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req schema.CreateUserReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}
	
	token, user, err := h.UserUC.Create(req.UserName, req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(500, "Internal server error")
	}

	res := schema.CreateUserRes{
		Token: token,
		User: schema.User{
			UserID: user.UserID,
			UserName: user.UserName,
			Email: user.Email,
			DisplayName: user.DisplayName,
			IconUrl: user.IconUrl,
		},
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req schema.LoginReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}

	token, user, err := h.UserUC.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(500, "Internal server error")
	}

	res := schema.LoginRes{
		Token: token,
		User: schema.User{
			UserID: user.UserID,
			UserName: user.UserName,
			Email: user.Email,
			DisplayName: user.DisplayName,
			IconUrl: user.IconUrl,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) FindMe(c echo.Context) error {
	user, err := middleware.GetUserFromContext(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, "Internal server error")
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