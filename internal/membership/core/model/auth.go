package model

import "time"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthToken struct {
	AccessToken  SigningToken `json:"access_token"`
	RefreshToken SigningToken `json:"refresh_token"`
}

type Auth struct {
	Secret string
	Expiry time.Duration
}

type SigningToken struct {
	Aud   string   `json:"aud"`
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Jti   string   `json:"jti"`
	Roles []string `json:"roles;omitempty"`
	Exp   int      `json:"exp"`
}

type AuthTokenData struct {
	Aud   string   `json:"aud"`
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Jti   string   `json:"jti"`
	Roles []string `json:"roles;omitempty"`
	Exp   int      `json:"exp"`
}

type AuthTokenResponse struct {
	AccessToken  AuthTokenData
	RefreshToken AuthTokenData
}
