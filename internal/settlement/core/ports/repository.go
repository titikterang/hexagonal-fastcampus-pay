package ports

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type SettlementRepository interface {
	// to datastore
	SaveSettlementReport(ctx context.Context, data model.SettlementPayload) error
	GetSettlementReport(ctx context.Context, accountNo, date string) ([]model.SettlementPayload, error)

	// idempotence
	EventIDExists(ctx context.Context, id string) bool
	SaveEventID(ctx context.Context, id string) error

	// update balance - money service
	UpdateUserBalance(ctx context.Context, accountNo string, amount decimal.Decimal) error

	// get trx info
	GetTransactionInfoByID(ctx context.Context, id string, settleType model.SettlementType) (types.TransactionInfoResult, error)
}
