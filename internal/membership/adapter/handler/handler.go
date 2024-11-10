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
	_, err := h.membershipService.SubmitLogin(ctx, model.LoginInfo{
		Username: request.GetUsername(),
		Password: request.GetPassword(),
	})
	return nil, err
}

func (h Handler) SubmitLogout(ctx context.Context, request *membership.LogoutRequest) (*membership.LogoutResponse, error) {
	//TODO implement me
	return nil, nil

}

func (h Handler) SubmitRegistration(ctx context.Context, request *membership.RegistrationRequest) (*membership.RegistrationResponse, error) {
	//TODO implement me
	return nil, nil
}
