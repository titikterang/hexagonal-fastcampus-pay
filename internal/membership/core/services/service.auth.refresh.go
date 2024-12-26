package services

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
)

func (s *MembershipService) RefreshToken(ctx context.Context, token string) (model.LoginResponse, error) {
	// validate incoming refresh token
	data, err := s.ValidateRSAToken(s.refreshKeypair.pubKey, token)
	if err != nil {
		log.Error().Msgf("failed to ValidateRSAToken, err %#v", err)
		return model.LoginResponse{
			Message: "failed to refresh user token, please re login",
		}, err
	}

	userInfo := data.(model.UserAuthInfo)
	// if valid cek from redis
	existingToken, err := s.repository.GetUserSessionFromCache(ctx, userInfo.AccountNumber)
	if err != nil {
		log.Error().Msgf("failed to GetUserSessionFromCache, err %#v", err)
		return model.LoginResponse{
			Message: "failed to refresh user token, please re login",
		}, err
	}

	// if from redis is not exists --> return err
	if existingToken != token {
		return model.LoginResponse{
			Message: "invalid refresh token, please relogin",
		}, err
	}

	newToken, refresh, err := s.constructToken(ctx, s.refreshKeypair.privKey, userInfo, s.config.Token.RefreshExpiry)
	if err != nil {
		return model.LoginResponse{
				Message: err.Error(),
			},
			errors.New("failed to refresh token")
	}

	return model.LoginResponse{
		Success:      true,
		Token:        newToken,
		RefreshToken: refresh,
		Message:      "refresh token successful",
	}, nil
}
