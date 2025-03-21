package input_port

import (
	"chat_server/domain/entity"
)

type CreateUserArgs struct {
	UserName string
	Email string
	Password string
}

type UpdateUserArgs struct {
	UserID string
	DisplayName string
	IconUrl string
}

type IUserUsecase interface {
	Authenticate(token string) (string, error) 
	FindByID(userID string) (entity.User, error)
	Create(args CreateUserArgs) error
	Login(email string, password string) (string, error)
	Update(args UpdateUserArgs) (entity.User, error)
}