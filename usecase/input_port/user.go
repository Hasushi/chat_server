package input_port

import "chat_server/domain/entity"

type IUserUsecase interface {
	Authenticate(token string) (string, error) 
	FindByID(userID string) (entity.User, error)
}