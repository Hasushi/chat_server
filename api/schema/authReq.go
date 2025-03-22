package schema

const Bearer = "Bearer"

type CreateUserReq struct {
	Email   string  `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

