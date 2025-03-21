package schema

type ProfileUpdateReq struct {
	Bio string `json:"bio"`
	IconUrl string `json:"iconUrl"`
}