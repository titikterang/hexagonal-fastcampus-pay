package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
)

type DatastoreRepositoryAdapter interface {
	// to redis
	GetUserSessionFromCache(ctx context.Context, accountNo string) (string, error)
	UpdateUserSessionIntoCache(ctx context.Context, accountNo, refreshData string) error
	DeleteUserSessionFromCache(ctx context.Context, accountNo string) error

	// cek user info from redis
	GetUserInfoFromCache(ctx context.Context, accountNumber string) (model.UserProfileInfo, error)
	UpdateUserInfoOnCache(ctx context.Context, data model.UserProfileInfo) error

	// to postgres
	InsertUserInfoIntoDB(ctx context.Context, payload model.RegistrationPayload) error
	GetUserInfoFromDB(ctx context.Context, accountNumber string) (model.UserProfileInfo, error)
	GetUserByUsername(ctx context.Context, username string) (model.UserAuthInfo, error)
}
