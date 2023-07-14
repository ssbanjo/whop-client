package types

type Customer struct {
	Id             string          `json:"id"`
	Username       string          `json:"username"`
	Email          string          `json:"email"`
	ProfilePicUrl  string          `json:"profile_pic_url"`
	SocialAccounts []SocialAccount `json:"social_accounts"`
	Roles          string          `json:"roles"`
}
