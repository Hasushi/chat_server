package schema

type CreateUserReq struct {
	UserName string `json:"username"`
	Email   string  `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

