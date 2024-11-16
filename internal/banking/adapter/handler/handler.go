package handler

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/banking"
)

func (h *Handler) SubmitDeposit(ctx context.Context, payload *banking.DepositRequest) (*banking.DepositResponse, error) {
	err := h.bankingService.SubmitDeposit(ctx, model.BankingCashHistory{
		DestinationAccount: payload.GetAccountNumber(),
		Amount:             float64(payload.GetAmount()),
		CashMovementType:   model.MovementDeposit,
	})
	return &banking.DepositResponse{
		Message: "deposit done",
		Status:  err == nil,
	}, err
}

func (h *Handler) SubmitPayment(ctx context.Context, payload *banking.BankPaymentRequest) (*banking.BankPaymentResponse, error) {
	err := h.bankingService.SubmitPayment(ctx, model.BankingCashHistory{
		SourceAccount:      payload.GetAccountNumberSource(),
		DestinationAccount: payload.GetAccountNumberDestination(),
		Amount:             float64(payload.GetAmount()),
		CashMovementType:   model.MovementPayment,
	})

	return &banking.BankPaymentResponse{
		Message:   "Payment Done",
		Status:    err == nil,
		PaymentId: payload.GetPaymentId(),
	}, err
}
func (h *Handler) SubmitTransfer(ctx context.Context, payload *banking.BankTransferRequest) (*banking.BankTransferResponse, error) {
	err := h.bankingService.SubmitTransfer(ctx, model.BankingCashHistory{
		SourceAccount:      payload.GetAccountNumberSource(),
		DestinationAccount: payload.GetAccountNumberDestination(),
		Amount:             float64(payload.GetAmount()),
		CashMovementType:   model.MovementTransfer,
	})
	return &banking.BankTransferResponse{
		Message: "Transfer Done",
		Status:  err == nil,
	}, err
}
