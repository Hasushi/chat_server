package database

import (
	"chat_server/domain/entity"
	"chat_server/usecase/output_port"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) output_port.User {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindByID(userID string) (entity.User, error) {
	return entity.User{}, nil
}

func (u *UserRepository) Create() error {
	return nil
}