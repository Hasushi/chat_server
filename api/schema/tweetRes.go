package schema

type Tweet struct {
	TweetID string `json:"tweetId"`
	Content string `json:"content"`
	ImageUrl string `json:"imageUrl"`
	CreatedAt string `json:"createdAt"`
	Author PublicUser `json:"author"`
	Likes []Like `json:"likes"`
	Retweets []Retweet `json:"retweets"`
	Replies []Replies `json:"replies"`
}

type Like struct {
	UserID string `json:"userId"`
}

type Retweet struct {
	UserID string `json:"userId"`
}

type Replies struct {
	ReplyID string `json:"replyId"`
}