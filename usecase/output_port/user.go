package outputport

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
}

type User interface {
	Create() error
	FindByID(userID string) (entity.User, error)
}