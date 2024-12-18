package ports

import (
	"context"
	"github.com/shopspring/decimal"
)

type SettlementRepository interface {
	// to datastore
	SaveSettlementReport(ctx context.Context) error
	GetSettlementReport(ctx context.Context, accountNo, date string)

	// idempotence
	EventIDExists(ctx context.Context, eventType string, accountNo, id string) bool
	SaveEventID(ctx context.Context, eventType string, accountNo, id string) error

	// update balance - money service
	UpdateUserBalance(ctx context.Context, accountNo string, amount decimal.Decimal) error
}
