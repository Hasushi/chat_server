package database

import (
	"chat_server/adapter/database/model"
	"chat_server/domain/entity"
	"chat_server/usecase/output_port"

	"gorm.io/gorm"
)

type TweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) output_port.Tweet {
	return &TweetRepository{db: db}
}

func (t *TweetRepository) Create(args output_port.CreateTweetArgs) error {
	return createTweet(t.db, args)
}

func (t *TweetRepository) CreateWithTx(tx interface{}, args output_port.CreateTweetArgs) error {
	db, ok := tx.(*gorm.DB)
	if !ok {
		return output_port.ErrInvalidTransaction
	}

	return createTweet(db, args)
}

func createTweet(db *gorm.DB, args output_port.CreateTweetArgs) error {
	tweet := model.Tweet{
		TweetID:  args.TweetID,
		UserID:   args.UserID,
		Content:  args.Content,
		ImageUrl: args.ImageUrl,
		CreatedAt: args.CreatedAt,
	}
	if err := db.Create(&tweet).Error; err != nil {
		return err
	}
	return nil
}

func (t *TweetRepository) Search(userID string, skip int, limit int) ([]entity.Tweet, error) {
	// TODO: Implement this
	// フォローしてる人だけのツイートを取得する
	return []entity.Tweet{}, nil
}

func (t *TweetRepository) Delete(userID string, tweetID string) error {
	// ツイートが存在するか確認
	var tweet model.Tweet
	if err := t.db.Where("tweet_id = ? AND user_id = ?", tweetID, userID).First(&tweet).Error; err != nil {
		return err
	}

	// ツイートを削除
	if err := t.db.Delete(&tweet).Error; err != nil {
		return err
	}
	return nil
}
