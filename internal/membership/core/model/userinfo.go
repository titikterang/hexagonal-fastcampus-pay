package model

import "time"

const (
	DefaultUserInfoTTL = 5 * time.Second
)

// UserAuthInfo - used by secondary ports
type UserAuthInfo struct {
	AccountNumber string    `json:"account_number" db:"account_number"`
	Username      string    `json:"username" db:"username"`
	Hash          string    `json:"hash" db:"hash"`
	LastLogin     time.Time `json:"last_login" db:"last_login"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// UserProfileInfo - to get user profile info
type UserProfileInfo struct {
	AccountNumber string `json:"account_number" db:"account_number"`
	Email         string `json:"email" db:"email"`
	Fullname      string `json:"fullname" db:"fullname"`
	Status        string `json:"status" db:"status"`
}

type RegistrationPayload struct {
	LoginInfo
	UserProfileInfo
	Hash string
}
