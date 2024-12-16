package handler

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"time"
)

func (h *Handler) GetPaymentStatus(ctx context.Context, payload *payment.PaymentStatusPayload) (*payment.PaymentStatusResponse, error) {
	data, err := h.paymentService.GetPaymentInfoByID(ctx, payload.GetInvoiceId())
	paymentHist := make([]*payment.PaymentInfo, 0)
	for _, v := range data.History {
		paymentHist = append(paymentHist, &payment.PaymentInfo{
			Datetime: v.DateTime,
			Status:   data.Info.GetStatusProto(),
		})
	}
	return &payment.PaymentStatusResponse{
		InvoiceId:      data.Info.InvoiceID,
		Status:         data.Info.GetStatusProto(),
		Amount:         data.Info.Amount.IntPart(),
		Datetime:       data.Info.DateTime.Format(time.DateTime),
		Message:        data.Info.Message,
		PaymentHistory: paymentHist,
	}, err
}

func (h *Handler) SubmitPayment(ctx context.Context, payload *payment.SubmitPaymentPayload) (*payment.SubmitPaymentResponse, error) {
	var (
		msg    string
		status = payment.PaymentStatus_PENDING
	)
	invoiceID, err := h.paymentService.SubmitPaymentRequest(ctx, model.SubmitPaymentPayload{
		MerchantID: payload.GetMerchantId(),
		Amount:     payload.GetAmount(),
		AccountNo:  payload.GetAccountNo(),
	})
	if err != nil {
		msg = err.Error()
		status = payment.PaymentStatus_FAILED
	}
	return &payment.SubmitPaymentResponse{
		InvoiceId: invoiceID,
		Status:    status,
		Message:   msg,
	}, nil
}
