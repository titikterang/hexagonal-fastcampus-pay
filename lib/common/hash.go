package common

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func GetHashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Stack().Err(err)
	}

	return string(hash)
}

func VerifyHash(hashedPwd string, plainTxt string) (valid bool) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainTxt))
	if err != nil {
		log.Error().Stack().Err(err)
		return
	}

	return true
}
