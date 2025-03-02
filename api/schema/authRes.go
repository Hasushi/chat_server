package schema
type CreateUserRes struct {
	UserID       string    `json:"userid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type LoginRes struct {
	Token string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}