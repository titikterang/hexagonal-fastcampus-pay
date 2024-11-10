package services

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
)

func (s *MembershipService) SubmitRegisterUser(ctx context.Context) {

}

func (s *MembershipService) GetUserInfo(ctx context.Context, accountNumber string) (model.UserProfileInfo, error) {
	return s.repository.GetUserInfoFromDB(ctx, accountNumber)
}

func (s *MembershipService) SubmitLogin(ctx context.Context, payload model.LoginInfo) (model.LoginResponse, error) {
	return model.LoginResponse{}, nil
}

func (s *MembershipService) SubmitLogout(ctx context.Context) {

}
