package router

import (
	"chat_server/api/handler"
	input_port "chat_server/usecase/input_port"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(
	userUC input_port.IUserUsecase,
) *http.Server {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))

	userHandler := handler.NewUserHandler(userUC)

	api := e.Group("/api/v1")
	auth := api.Group("/auth")
	auth.POST("/register", userHandler.CreateUser)
	auth.POST("/login", userHandler.Login)

	return e.Server
}