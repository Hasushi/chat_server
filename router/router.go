package router

import (
	"chat_server/api/handler"
	input_port "chat_server/usecase/input_port"
	"net/http"
	apiMiddleware "chat_server/middleware"

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

	authHandler := handler.NewAuthHandler(userUC)
	userHandler := handler.NewUserHandler(userUC)

	api := e.Group("/api/v1")
	auth := api.Group("/auth")
	auth.POST("/register", authHandler.CreateUser)
	auth.POST("/login", authHandler.Login)
	
	users := api.Group("/users", apiMiddleware.NewAuthMiddleware(userUC).Authenticate)
	users.GET("/me", userHandler.FindMe)

	return e.Server
}