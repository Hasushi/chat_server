package schema

type User struct {
	UserID       string    `json:"userid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserRes struct {
	User User `json:"user"`
	Token string `json:"token"`
}

type LoginRes struct {
	Token string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}