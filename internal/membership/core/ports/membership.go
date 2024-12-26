package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
)

type MembershipServiceAdapter interface {
	SubmitRegisterUser(ctx context.Context, payload model.RegistrationPayload) (string, error)
	GetUserInfo(ctx context.Context, accountNumber string) (model.UserProfileInfo, error)
	SubmitLogin(ctx context.Context, payload model.LoginInfo) (model.LoginResponse, error)
	RefreshToken(ctx context.Context, token string) (model.LoginResponse, error)
	SubmitLogout(ctx context.Context, uuid string) error
}
