package input_port

type CreateTweetArgs struct {
	UserID string
	Content string
	ImageUrl string
}

type ITweetUsecase interface {
	Create(args CreateTweetArgs) error
}
