//go:generate mockgen -source=repository.go -destination=mocks/mock.repository.go -package=mocks
package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type MoneyRepositoryAdapter interface {
	// method to redis cache
	GetSnapshot(ctx context.Context, accountNumber string) (string, error)
	UpdateSnapshot(ctx context.Context, accountNumber, amount string) error
	GetCashMovementFromCache(ctx context.Context, accountNumber string) ([]model.CashMovementInfo, error)
	AppendCashMovementInfoIntoCache(ctx context.Context, info model.CashMovementInfo) error

	// prevent double update
	ReqIDExists(ctx context.Context, accountNumber, reqID string) bool
	SaveReqID(ctx context.Context, accountNumber, reqID string)

	// method to db
	AppendCashMovementIntoDatastore(ctx context.Context, info model.CashMovementInfo) error
	ConstructBalanceInfo(ctx context.Context, info model.UserCashInfo) error
	GetBalanceInfoFromDB(ctx context.Context, accountNo string) (model.UserCashInfo, error)

	// kafka producer
	PublishMoneyValidationMessageReply(ctx context.Context, info types.TransactionValidateReplyInfo) error
}
