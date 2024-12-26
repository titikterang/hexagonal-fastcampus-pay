package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
	"io"
	"time"
)

func (s *MembershipService) CreateToken(user string) (model.Token, error) {
	claims := jwt.MapClaims{}
	claims["user"] = user
	claims["exp"] = time.Now().Add(s.config.Token.Expiry).Unix()
	jwtModel := model.Token{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var err error
	jwtModel.AccessToken, err = token.SignedString([]byte(s.config.Token.Secret))
	if err != nil {
		return jwtModel, err
	}

	return s.Refresh(jwtModel)
}

func (s *MembershipService) Validate(token model.Token) (string, error) {
	sha1Var := sha1.New()
	io.WriteString(sha1Var, s.config.Token.Secret)

	var username string
	salt := string(sha1Var.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return username, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return username, err
	}

	var data string
	_, err = base64.URLEncoding.Decode([]byte(data), []byte(token.RefreshToken))
	if err != nil {
		return username, err
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plain, err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		return username, err
	}

	if string(plain) != token.AccessToken {
		return username, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	newToken, _, err := parser.ParseUnverified(token.AccessToken, claims)
	if err != nil {
		return username, err
	}

	payload, ok := newToken.Claims.(jwt.MapClaims)
	if !ok {
		return username, errors.New("invalid token")
	}

	username = payload["user"].(string)
	return username, nil
}

func (s *MembershipService) Refresh(token model.Token) (model.Token, error) {
	sha1Var := sha1.New()
	io.WriteString(sha1Var, s.config.Token.Secret)

	salt := string(sha1Var.Sum(nil)[0:16])
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return model.Token{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return model.Token{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return model.Token{}, err
	}

	token.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(token.AccessToken), nil))
	return token, nil
}
