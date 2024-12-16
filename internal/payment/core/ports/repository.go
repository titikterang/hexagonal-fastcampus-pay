package ports

import (
	"context"
	model2 "github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type PaymentRepositoryAdapter interface {
	// save to db
	InsertPaymentHistory(ctx context.Context, data model2.PaymentInfo)
	UpdatePaymentStatus(ctx context.Context, data model2.PaymentInfo)
	InsertPaymentInfo(ctx context.Context)
	GetPaymentInfo(ctx context.Context, invoiceID string) (model2.PaymentInfo, []model2.PaymentHistory, error)

	// idempotence check
	EventIDExists(ctx context.Context, eventType string, accountNo, id string) bool
	SaveEventID(ctx context.Context, eventType string, accountNo, id string) error

	// external api call
	GetAccountInfo(ctx context.Context, accountNumber string) (types.UserProfileInfo, error)

	// publish to other service
	PublishPaymentValidateRequest(ctx context.Context, info types.TransactionValidateInfo) error // to money service
	PublishPaymentBankingRequest(ctx context.Context, payload types.PaymentBankExecution) error  // to banking service
}
