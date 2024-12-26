package handler

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/membership"
)

func (h Handler) GetUserInfo(ctx context.Context, payload *membership.UserInfoPayload) (*membership.UserInfoResponse, error) {
	data, err := h.membershipService.GetUserInfo(ctx, payload.GetAccountNumber())
	if err != nil {
		return nil, err
	}
	return &membership.UserInfoResponse{
		Email:         data.Email,
		Fullname:      data.Fullname,
		AccountNumber: data.AccountNumber,
		Status:        data.Status,
	}, nil
}

func (h Handler) SubmitLogin(ctx context.Context, request *membership.LoginRequest) (*membership.LoginResponse, error) {
	resp, err := h.membershipService.SubmitLogin(ctx, model.LoginInfo{
		Username: request.GetUsername(),
		Password: request.GetPassword(),
	})
	return &membership.LoginResponse{
		Success:      resp.Success,
		LoginMessage: resp.Message,
		Token:        resp.Token,
		RefreshToken: resp.RefreshToken,
	}, err
}

func (h Handler) SubmitLogout(ctx context.Context, request *membership.LogoutRequest) (*membership.LogoutResponse, error) {
	err := h.membershipService.SubmitLogout(ctx, request.GetUuid())
	return &membership.LogoutResponse{
		Message: "Logout Success",
	}, err
}

func (h Handler) SubmitRegistration(ctx context.Context, request *membership.RegistrationRequest) (*membership.RegistrationResponse, error) {
	accno, err := h.membershipService.SubmitRegisterUser(ctx, model.RegistrationPayload{
		LoginInfo: model.LoginInfo{
			Username: request.GetUsername(),
			Password: request.GetPassword(),
		},
		UserProfileInfo: model.UserProfileInfo{
			Email:    request.GetEmail(),
			Fullname: request.GetFullname(),
		},
	})
	return &membership.RegistrationResponse{
		AccountNumber: accno,
		Success:       err == nil,
		ErrorMessage: func() string {
			if err != nil {
				return err.Error()
			}

			return ""
		}(),
	}, err
}

func (h Handler) RefreshToken(ctx context.Context, payload *membership.RefreshRequest) (*membership.LoginResponse, error) {
	data, err := h.membershipService.RefreshToken(ctx, payload.GetRefreshToken())
	return &membership.LoginResponse{
		Success:      data.Success,
		LoginMessage: data.Message,
		Token:        data.Token,
		RefreshToken: data.RefreshToken,
	}, err
}
