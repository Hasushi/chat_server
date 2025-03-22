package schema

type CreateTweetReq struct {
	Content string `json:"content"`
	// TODO いずれ何枚かの画像をアップロードできるようにする
	ImageUrl string `json:"imageUrl"`
}