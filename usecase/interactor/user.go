package interactor
import (
	output_port"chat_server/usecase/output_port"
	input_port"chat_server/usecase/input_port"
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

func (u *UserUsecase) Create() {
}

func (u *UserUsecase) Login() {
}