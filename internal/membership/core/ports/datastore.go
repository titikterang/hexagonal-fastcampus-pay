package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
)

type DatastoreRepositoryAdapter interface {
	// to redis
	GetUserSessionFromCache(ctx context.Context)
	UpdateUserSessionIntoCache(ctx context.Context)

	// to postgres
	InsertUserInfoIntoDB(ctx context.Context, payload model.RegistrationPayload) error
	GetUserInfoFromDB(ctx context.Context, accountNumber string) (model.UserProfileInfo, error)
	GetUserByUsername(ctx context.Context)
}
