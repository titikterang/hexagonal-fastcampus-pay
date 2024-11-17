package services

import (
	"context"
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

func (s *MembershipService) SubmitLogin(ctx context.Context, payload model.LoginInfo) (model.LoginResponse, error) {
	return model.LoginResponse{}, nil
}

func (s *MembershipService) SubmitLogout(ctx context.Context, uuid string) error {

	return nil
}
