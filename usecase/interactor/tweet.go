package interactor

import (
	"chat_server/usecase/input_port"
	"chat_server/usecase/output_port"
)

type TweetUsecase struct {
	tweetRepo output_port.Tweet
	clock    output_port.Clock
	ulid    output_port.ULID
	transaction output_port.GormTransaction
}

type NewTweetUsecaseArgs struct {	
	TweetRepo output_port.Tweet
	Clock     output_port.Clock
	ULID      output_port.ULID
	Transaction output_port.GormTransaction
}

func NewTweetUsecase(args NewTweetUsecaseArgs) input_port.ITweetUsecase {
	return &TweetUsecase{
		tweetRepo: args.TweetRepo,
		clock: args.Clock,
		ulid: args.ULID,
		transaction: args.Transaction,
	}
}

func (u *TweetUsecase) Create(args input_port.CreateTweetArgs) error {
	tweetID := u.ulid.GenerateID()
	err := u.transaction.StartTransaction(func(tx interface{}) error {
		err := u.tweetRepo.CreateWithTx(tx, output_port.CreateTweetArgs{
			TweetID: tweetID,
			UserID: args.UserID,
			Content: args.Content,
			ImageUrl: args.ImageUrl,
			CreatedAt: u.clock.Now(),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
