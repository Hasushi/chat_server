package model

import (
	"chat_server/domain/entity"
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	TweetID  string
	UserID   string
	Content  string
	ImageUrl string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t Tweet) ToEntity() entity.Tweet {
	return entity.Tweet{
		TweetID:  t.TweetID,
		UserID:   t.UserID,
		Content:  t.Content,
		ImageUrl: t.ImageUrl,
	}
}