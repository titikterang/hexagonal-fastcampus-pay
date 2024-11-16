package services

import (
	"context"
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
	return s.repository.GetUserInfoFromDB(ctx, accountNumber)
}

func (s *MembershipService) SubmitLogin(ctx context.Context, payload model.LoginInfo) (model.LoginResponse, error) {
	return model.LoginResponse{}, nil
}

func (s *MembershipService) SubmitLogout(ctx context.Context, uuid string) error {
	return nil
}
