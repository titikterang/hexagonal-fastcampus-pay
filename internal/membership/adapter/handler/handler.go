package handler

import (
	"context"
	"errors"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/common"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/membership"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/tracer"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h Handler) HealthCheck(context.Context, *emptypb.Empty) (*membership.HealthResponse, error) {
	return &membership.HealthResponse{
		Response: "ok",
	}, nil
}

func (h Handler) GetUserInfo(ctx context.Context, payload *membership.UserInfoPayload) (*membership.UserInfoResponse, error) {
	ctx, span := tracer.StartInitSpan(ctx, "GetUserInfo")
	defer span.End()
	data, err := h.membershipService.GetUserInfo(ctx, payload.GetAccountNumber())
	if err != nil {
		span.RecordError(err)
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
	ctx, span := tracer.StartInitSpan(ctx, "SubmitLogin")
	defer span.End()

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

func (h Handler) SubmitLogout(ctx context.Context, _ *emptypb.Empty) (*membership.LogoutResponse, error) {
	userID, ok := common.ExtractUserIDFromHeader(ctx)
	ctx, span := tracer.StartInitSpan(ctx, "SubmitLogout")
	defer span.End()

	if !ok {
		return &membership.LogoutResponse{
			Message: "Logout Success",
		}, errors.New("invalid user ID")
	}
	err := h.membershipService.SubmitLogout(ctx, userID)
	if err != nil {
		return &membership.LogoutResponse{
			Message: "Failed to logout",
		}, err
	}
	return &membership.LogoutResponse{
		Message: "Logout Success",
	}, nil
}

func (h Handler) SubmitRegistration(ctx context.Context, request *membership.RegistrationRequest) (*membership.RegistrationResponse, error) {
	ctx, span := tracer.StartInitSpan(ctx, "SubmitRegistration")
	defer span.End()

	//_, ok := common.ExtractUserIDFromHeader(ctx)
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
	ctx, span := tracer.StartInitSpan(ctx, "RefreshToken")
	defer span.End()

	//_, ok := common.ExtractUserIDFromHeader(ctx)
	data, err := h.membershipService.RefreshToken(ctx, payload.GetRefreshToken())
	return &membership.LoginResponse{
		Success:      data.Success,
		LoginMessage: data.Message,
		Token:        data.Token,
		RefreshToken: data.RefreshToken,
	}, err
}
