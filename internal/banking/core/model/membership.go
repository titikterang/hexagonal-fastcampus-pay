package model

const (
	UserStatusActive  = "A"
	UserStatusSuspend = "S"
)

// UserProfileInfo - to get user profile info
type UserProfileInfo struct {
	AccountNumber string `json:"account_number"`
	Email         string `json:"email"`
	Fullname      string `json:"fullname"`
	Status        string `json:"status"`
}

type UserProfileAPIResult struct {
	Email         string `json:"email"`
	Fullname      string `json:"fullname"`
	AccountNumber string `json:"accountNumber"`
	Status        string `json:"status"`
}
