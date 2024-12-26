package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
	"time"
)

func (s *MembershipService) CreateRSAToken(userInfo model.UserAuthInfo) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.privKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["dat"] = userInfo                              // Our custom data.
	claims["exp"] = now.Add(s.config.Token.Expiry).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()                            // The time at which the token was issued.
	claims["nbf"] = now.Unix()                            // The time before which the token must be disregarded.

	tokenData := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenData.Header["kid"] = s.config.Token.KeyID
	token, err := tokenData.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	//return base64.URLEncoding.EncodeToString([]byte(token)), nil
	return token, nil
}

func (s *MembershipService) ValidateRSAToken(token string) (interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(s.pubKey)
	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims["dat"], nil
}
