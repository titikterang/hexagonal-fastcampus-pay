package services

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/common"
	"strconv"
	"time"
)

func (s *MembershipService) SubmitRegisterUser(ctx context.Context, payload model.RegistrationPayload) (string, error) {
	// generate account number
	payload.AccountNumber = strconv.FormatInt(time.Now().UnixMilli(), 10)

	// get hash from password
	payload.Hash = common.GetHashAndSalt(payload.Password)

	// insert info
	return payload.AccountNumber, s.repository.InsertUserInfoIntoDB(ctx, payload)
}

func (s *MembershipService) GetUserInfo(ctx context.Context, accountNumber string) (model.UserProfileInfo, error) {
	// cek from redis
	// if exists, then return from redis
	log.Info().Msg("get user info ....")
	data, err := s.repository.GetUserInfoFromCache(ctx, accountNumber)
	if err == nil {
		log.Info().Msg("return from redis ....")
		return data, nil
	}

	if err != nil {
		// no exists
		log.Err(err).Msg("GetUserInfo.GetUserInfoFromCache")
	}

	// if not exists -> query from DB, save to redis with TTL 1hour
	dbUserInfo, err := s.repository.GetUserInfoFromDB(ctx, accountNumber)
	if err != nil {
		log.Err(err).Msg("GetUserInfo.GetUserInfoFromDB")
		return dbUserInfo, err
	}

	// update cache info
	err = s.repository.UpdateUserInfoOnCache(ctx, dbUserInfo)
	if err != nil {
		log.Err(err).Msg("GetUserInfo.UpdateUserInfoOnCache")
	}
	log.Info().Msg("return from db ....")
	return dbUserInfo, err
}

func (s *MembershipService) constructToken(ctx context.Context, key []byte, userInfo model.UserAuthInfo, expiry time.Duration) (token string, refresh string, err error) {
	// generate auth token
	data, err := s.CreateRSAToken(key, expiry, userInfo)
	if err != nil {
		return data, "", err
	}

	// generate refresh token
	refreshData, err := s.CreateRSAToken(s.refreshKeypair.privKey, s.config.Token.RefreshExpiry, userInfo)
	if err != nil {
		return data, refreshData, err
	}

	// save to redis
	// se ttl based on refresh expiry
	err = s.repository.UpdateUserSessionIntoCache(ctx, userInfo.AccountNumber, refreshData)
	return data, refreshData, err
}

func (s *MembershipService) SubmitLogin(ctx context.Context, payload model.LoginInfo) (model.LoginResponse, error) {
	// query db, get user info by DB
	userInfo, err := s.repository.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return model.LoginResponse{}, err
	}

	if !common.VerifyHash(userInfo.Hash, payload.Password) {
		return model.LoginResponse{}, errors.New("invalid username & password combination")
	}

	token, refresh, err := s.constructToken(ctx, s.authKeyPair.privKey, userInfo, s.config.Token.Expiry)
	if err != nil {
		return model.LoginResponse{
				Message: err.Error(),
			},
			errors.New("failed to generate token")
	}
	return model.LoginResponse{
		Success:      true,
		Token:        token,
		RefreshToken: refresh,
		Message:      "login successful",
	}, nil
}

func (s *MembershipService) SubmitLogout(ctx context.Context, accountNumber string) error {
	return s.repository.DeleteUserSessionFromCache(ctx, accountNumber)
}
