package authentication

import (
	output_port "chat_server/usecase/output_port"
	"time"
)

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

func (u *UserAuth) CheckPassword(hashedPassword, password string) error {
	return CheckBcryptPassword(hashedPassword, password)
}

func (u *UserAuth) IssueUserToken(userID string, issuedAt time.Time) (string, error) {
	return IssueUserToken(userID, issuedAt, []string{output_port.TokenScopeGeneral})
}