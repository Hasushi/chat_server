package handler

import (
	"chat_server/api/schema"
	input_port "chat_server/usecase/input_port"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	UserUC input_port.IUserUsecase
}

func NewAuthHandler(userUC input_port.IUserUsecase) *AuthHandler {
	return &AuthHandler{
		UserUC: userUC,
	}
}

func (h *AuthHandler) CreateUser(c echo.Context) error {
	var req schema.CreateUserReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	
	err := h.UserUC.Create(input_port.CreateUserArgs{
		UserName: req.UserName,
		Email: req.Email,
		Password: req.Password,
	})
	
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.NoContent(http.StatusCreated)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req schema.LoginReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	token, err := h.UserUC.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	res := schema.LoginRes{
		Token: token,
	}

	return c.JSON(http.StatusOK, res)
}
