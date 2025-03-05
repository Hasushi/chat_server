package input_port

import (
	"chat_server/domain/entity"
	"time"
)

type IUserUsecase interface {
	Authenticate(token string) (string, error) 
	FindByID(userID string) (entity.User, error)
	Create(userName string, email string, password string) (entity.User, error)
	Login(email string, password string) (string, time.Duration, error)
}