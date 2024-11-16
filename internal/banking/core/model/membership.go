package model

// UserProfileInfo - to get user profile info
type UserProfileInfo struct {
	AccountNumber string `json:"account_number"`
	Email         string `json:"email"`
	Fullname      string `json:"fullname"`
	Status        string `json:"status"`
}
