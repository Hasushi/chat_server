package output_port

import (
	"chat_server/domain/entity"
	"time"
)

type CreateTweetArgs struct {
	TweetID string
	UserID string
	Content string
	ImageUrl string
	CreatedAt time.Time
}

type Tweet interface {
	Create(args CreateTweetArgs) error
	CreateWithTx(tx interface{}, args CreateTweetArgs) error
	Search(userID string, skip int, limit int) ([]entity.Tweet, error)
	Delete(userID string, tweetID string) error
}