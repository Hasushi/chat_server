package output_port

import (
	"chat_server/domain/entity"
	"time"
	"errors"
)

var (
	TokenScopeGeneral = "general"
	TokenGeneralExpireDuration = 7 * 24 * time.Hour
	ErrUnknownScope = errors.New("unknown scope")
	ErrTokenScopeInvalid = errors.New("token scope invalid")
)

type UserAuth interface {
	Authenticate(token string) (string, error)
	CheckPassword(hashedPassword, password string) error
	HashPassword(password string) (string, error)
	IssueUserToken(userID string, issuedAt time.Time) (string, error)
}

type CreateUserArgs struct {
	UserID string
	UserName string
	Email string
	HashedPassword string
	// TODO iconどうするか
}

type User interface {
	Create(args CreateUserArgs) error
	FindByID(userID string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}