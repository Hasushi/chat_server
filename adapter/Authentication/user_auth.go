package authentication

import output_port "chat_server/usecase/output_port"

type UserAuth struct {}

func NewUserAuth() output_port.UserAuth {
	return &UserAuth{}
}

func (u *UserAuth) Authenticate(token string) (string, error) {
	return "", nil
}

func (u *UserAuth) HashPassword(password string) (string, error) {
	hp, err := HashBcryptPassword(password)
	if err != nil {
		return "", err
	}
	return hp, nil
}