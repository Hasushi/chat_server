package authentication

import (
	"chat_server/config"
	output_port "chat_server/usecase/output_port"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Scopes []string
}

var userToken = &hs256jwt{
	sigKey: []byte(config.GetSigKey()),
	createClaims: func() jwt.Claims {
		return &UserClaims{}
	},
}

func IssueUserToken(userID string, issuedAt time.Time, scopes []string) (string, error) {
	id := uuid.New().String()

	var expiredAt time.Time
	if isContains(scopes, output_port.TokenScopeGeneral) {
		expiredAt = issuedAt.Add(output_port.TokenGeneralExpireDuration)
	} else {
		return "", output_port.ErrUnknownScope
	}

	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
		Scopes: scopes,
	}

	return userToken.issueToken(claims)
}

func VerifyUserToken(tokenString string, scopes []string) (string, error) {
	claims, err := userToken.parseToken(tokenString)
	if err != nil {
		return "", err
	}

	if !isInclusive(claims.(*UserClaims).Scopes, scopes) {
		return "", output_port.ErrTokenScopeInvalid
	}
	return claims.(*UserClaims).Subject, nil
}

func isInclusive(x, y []string) bool {
	for _, v := range y {
		if !isContains(x, v) {
			return false
		}
	}
	return true
}

func isContains(hasScopes []string, scope string) bool {
	for _, hasScope := range hasScopes {
		if hasScope == scope {
			return true
		}
	}
	return false
}