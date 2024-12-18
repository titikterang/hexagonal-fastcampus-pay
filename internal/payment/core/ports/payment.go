package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type PaymentServiceAdapter interface {
	SubmitPaymentRequest(ctx context.Context, payload model.SubmitPaymentPayload) (string, error)
	GetPaymentInfoByID(ctx context.Context, id string) (model.PaymentDetails, error)
	GetPaymentPrecheckInfo(ctx context.Context, accountNo string) model.PaymentPrecheckInfo
	GetInfoByTrxID(ctx context.Context, id string) (types.TransactionInfoResult, error)
}

type PaymentConsumerAdapter interface {
	HandleTransactionValidationReply(ctx context.Context, data types.TransactionValidateReplyInfo) error
	HandleBankCallbackReply(ctx context.Context, data types.PaymentBankReply) error
}
