package authentication

import output_port "chat_server/usecase/output_port"

type UserAuth struct {}

func NewUserAuth() output_port.UserAuth {
	return &UserAuth{}
}

func (u *UserAuth) Authenticate(token string) (string, error) {
	return "", nil
}