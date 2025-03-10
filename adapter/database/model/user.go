package model

import (
	"chat_server/domain/entity"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID       string    
	UserName string 
	Email    string 
	HashedPassword string
	DisplayName string
	IconUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt gorm.DeletedAt
}

func (u User) ToEntity() entity.User {
	return entity.User{
		UserID: u.UserID,
		UserName: u.UserName,
		Email: u.Email,
		HashedPassword: u.HashedPassword,
		DisplayName: u.DisplayName,
		IconUrl: u.IconUrl,
	}
}