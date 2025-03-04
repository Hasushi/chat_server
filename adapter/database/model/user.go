package model

import "chat_server/domain/entity"

type User struct {
	UserID       string    
	UserName string 
	Email    string 
	HashedPassword string 
}

func (u User) ToEntity() entity.User {
	return entity.User{
		UserID: u.UserID,
		UserName: u.UserName,
		Email: u.Email,
	}
}