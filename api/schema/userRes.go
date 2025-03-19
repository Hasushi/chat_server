package schema

type User struct {
	UserID       string `json:"userId"`
	Email   string `json:"email"`
	UserName string `json:"userName"`
	DisplayName string `json:"displayName"`
	IconUrl string `json:"iconUrl"`
}

type FindUserRes User

type ProfileUpdateReq struct {
	DisplayName string `json:"displayName"`
	IconUrl string `json:"iconUrl"`
}