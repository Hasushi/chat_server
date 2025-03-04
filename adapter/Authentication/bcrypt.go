package authentication

import (
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrHashPassword = errors.New("failed to hash password")
	ErrBase64Decode = errors.New("failed to decode base64")
	ErrPasswordMismatch = errors.New("password mismatch")
	ErrUnexpected = errors.New("unexpected error")
)

func bcryptHash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), nil
}

func HashBcryptPassword(password string) (string, error) {
	hash, err := bcryptHash(password)
	if err != nil {
		return "", ErrHashPassword
	}
	return hash, nil
}

func CheckPasswordHash(hashedPassword, password string) error {
	decodedHash, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return ErrBase64Decode
	}
	
	err = bcrypt.CompareHashAndPassword(decodedHash, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrPasswordMismatch
	} else if err != nil {
		return ErrUnexpected
	}
	return nil
}