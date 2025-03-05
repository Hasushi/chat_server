package interactor

import (
	"chat_server/domain/entity"
	input_port "chat_server/usecase/input_port"
	output_port "chat_server/usecase/output_port"
)

type UserUsecase struct {
	auth output_port.UserAuth
	user output_port.User
	output_port.ULID
}

type NewUserUsecaseArgs struct {
	Auth output_port.UserAuth
	User output_port.User
	ULID output_port.ULID
}

func NewUserUsecase(args NewUserUsecaseArgs) input_port.IUserUsecase {
	// TODO なんでここでポインタを返しているのかを理解する
	return &UserUsecase{
		auth: args.Auth,
		user: args.User,
		ULID: args.ULID,
	}
}

func (u *UserUsecase) Authenticate(token string) (string, error) {
	return u.auth.Authenticate(token)
}

func (u *UserUsecase) FindByID(userID string) (entity.User, error) {
	return entity.User{}, nil
}

func (u *UserUsecase) Create(userName string, email string, password string) (entity.User, error) {
	userID := u.ULID.GenerateID()
	hp, err := u.auth.HashPassword(password)
	if err != nil {
		return entity.User{}, err
	}

	args := output_port.CreateUserArgs{
		UserID: userID,
		UserName: userName,
		Email: email,
		HashedPassword: hp,
	}

	err = u.user.Create(args)
	if err != nil {
		return entity.User{}, err
	}

	user, err := u.user.FindByID(userID)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserUsecase) Login(email string, password string) (entity.User, error) {
	return entity.User{}, nil
}