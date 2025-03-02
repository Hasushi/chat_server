package interactor

import (
	"chat_server/domain/entity"
	input_port "chat_server/usecase/input_port"
	output_port "chat_server/usecase/output_port"
)

type UserUsecase struct {
	auth output_port.UserAuth
}

type NewUserUsecaseArgs struct {
	Auth output_port.UserAuth
}

func NewUserUsecase(args NewUserUsecaseArgs) input_port.IUserUsecase {
	// TODO なんでここでポインタを返しているのかを理解する
	return &UserUsecase{
		auth: args.Auth,
	}
}

func (u *UserUsecase) Authenticate(token string) (string, error) {
	return u.auth.Authenticate(token)
}

func (u *UserUsecase) FindByID(userID string) (entity.User, error) {
	return entity.User{}, nil
}

func (u *UserUsecase) Create(userName string, email string, password string) (entity.User, error) {
	return entity.User{}, nil
}

func (u *UserUsecase) Login(email string, password string) (entity.User, error) {
	return entity.User{}, nil
}