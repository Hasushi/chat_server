package schema


type CreateUserRes struct {
	Token string `json:"token"`
	User User `json:"user"`
}

type LoginRes struct {
	Token string `json:"token"`
	User User `json:"user"`
}