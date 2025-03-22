package entity

import "time"

type Tweet struct {
	TweetID  string
	UserID   string
	Content  string
	ImageUrl string
	CreatedAt time.Time
}