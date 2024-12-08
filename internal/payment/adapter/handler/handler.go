package handler

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
)

func (h *Handler) GetPaymentStatus(context.Context, *payment.PaymentStatusPayload) (*payment.PaymentStatusResponse, error) {
	return nil, nil
}

func (h *Handler) SubmitPayment(context.Context, *payment.SubmitPaymentPayload) (*payment.SubmitPaymentResponse, error) {
	return nil, nil
}
