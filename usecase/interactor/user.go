package interactor

import (
	"chat_server/domain/entity"
	input_port "chat_server/usecase/input_port"
	output_port "chat_server/usecase/output_port"
	"errors"
)

type UserUsecase struct {
	auth output_port.UserAuth
	user output_port.User
	output_port.ULID
	clock output_port.Clock
	transaction output_port.GormTransaction
}

type NewUserUsecaseArgs struct {
	Auth output_port.UserAuth
	User output_port.User
	ULID output_port.ULID
	Clock output_port.Clock
	Transaction output_port.GormTransaction
}

func NewUserUsecase(args NewUserUsecaseArgs) input_port.IUserUsecase {
	// TODO なんでここでポインタを返しているのかを理解する
	return &UserUsecase{
		auth: args.Auth,
		user: args.User,
		ULID: args.ULID,
		clock: args.Clock,
		transaction: args.Transaction,
	}
}

func (u *UserUsecase) Authenticate(token string) (string, error) {
	return u.auth.Authenticate(token)
}

func (u *UserUsecase) FindByID(userID string) (entity.User, error) {
	user, err := u.user.FindByID(userID)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserUsecase) Create(args input_port.CreateUserArgs) error {
	// アカウントが存在するかチェック
	_, err := u.user.FindByEmail(args.Email)
	if err == nil {
		return errors.New("email already exists")
	}
	
	userID := u.ULID.GenerateID()
	hp, err := u.auth.HashPassword(args.Password)
	if err != nil {
		return err
	}

	createArgs := output_port.CreateUserArgs{
		UserID: userID,
		UserName: args.UserName,
		Email: args.Email,
		HashedPassword: hp,
	}

	err = u.transaction.StartTransaction(func(tx interface{}) error {
		err := u.user.CreateWithTx(tx, createArgs)
		if err != nil {
			return err
		}

		_, err = u.user.FindByIDWithTx(tx, userID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Login(email string, password string) (string, error) {
	user, err := u.user.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = u.auth.CheckPassword(user.HashedPassword, password)
	if err != nil {
		return "", err
	}

	token, err := u.auth.IssueUserToken(user.UserID, u.clock.Now())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUsecase) Update(args input_port.UpdateUserArgs) (entity.User, error) {
	user, err := u.user.FindByID(args.UserID)
	if err != nil {
		return entity.User{}, err
	}

	updateArgs := output_port.UpdateUserArgs{
		UserID: user.UserID,
		DisplayName: args.DisplayName,
		IconUrl: args.IconUrl,
	}

	err = u.user.Update(updateArgs)
	if err != nil {
		return entity.User{}, err
	}

	updatedUser, err := u.user.FindByID(args.UserID)
	if err != nil {
		return entity.User{}, err
	}

	return updatedUser, nil
}