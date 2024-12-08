package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type PaymentRepositoryAdapter interface {
	// save to db
	InsertPaymentHistory(ctx context.Context)
	UpdatePaymentStatus(ctx context.Context)
	InsertPaymentInfo(ctx context.Context)

	// idempotence check
	EventIDExists(ctx context.Context, eventType model.EventType, accountNo, id string) bool
	SaveEventID(ctx context.Context, eventType model.EventType, accountNo, id string) error

	// external api call
	GetAccountInfo(ctx context.Context, accountNumber string) (types.UserProfileInfo, error)

	// publish to other service
	PublishPaymentValidateRequest(ctx context.Context, info types.TransactionValidateInfo) error // to money service
	PublishPaymentBankingRequest(ctx context.Context, payload types.PaymentBankExecution) error  // to banking service
}
