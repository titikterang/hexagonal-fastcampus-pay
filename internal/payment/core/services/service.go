package services

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

func (p *PaymentService) SubmitPaymentRequest(ctx context.Context, payload model.SubmitPaymentPayload) error {
	return nil
}
func (p *PaymentService) GetPaymentInfoByID(ctx context.Context, id string) (model.PaymentDetails, error) {
	return model.PaymentDetails{}, nil
}
func (p *PaymentService) HandleTransactionValidationReply(ctx context.Context, data types.TransactionValidateReplyInfo) error {
	return nil
}
func (p *PaymentService) HandleBankCallbackReply(ctx context.Context, data types.PaymentBankReply) error {
	return nil
}
