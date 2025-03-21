package schema

type User struct {
	UserID       string `json:"userId"`
	Email   string `json:"email"`
	UserName string `json:"userName"`
	Bio string `json:"bio"`
	IconUrl string `json:"iconUrl"`
}

type FindUserRes User
