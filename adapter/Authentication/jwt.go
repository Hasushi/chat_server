package authentication

import (
	"github.com/golang-jwt/jwt/v5"
)

type hs256jwt struct {
	sigKey []byte
	createClaims func() jwt.Claims
}

func (t *hs256jwt) issueToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.sigKey)
}

func (t *hs256jwt) parseToken(tokenString string) (jwt.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return t.sigKey, nil
	}
	token, err := jwt.ParseWithClaims(tokenString, t.createClaims(), keyFunc)
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}


