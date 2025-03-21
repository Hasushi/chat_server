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
	Bio string
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
		Bio: u.Bio,
		IconUrl: u.IconUrl,
	}
}