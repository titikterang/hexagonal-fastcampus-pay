package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type PaymentServiceAdapter interface {
	SubmitPaymentRequest(ctx context.Context, payload model.SubmitPaymentPayload) error
	GetPaymentInfoByID(ctx context.Context, id string) (model.PaymentDetails, error)
}

type PaymentConsumerAdapter interface {
	HandleTransactionValidationReply(ctx context.Context, data types.TransactionValidateReplyInfo) error
	HandleBankCallbackReply(ctx context.Context, data types.PaymentBankReply) error
}
