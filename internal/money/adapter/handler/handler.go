package handler

import (
	"context"
	"errors"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/common"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
)

// GetUserBalance - oublic endpoint , get from redis snapshoot, return amount (string)
func (h *Handler) GetUserBalance(ctx context.Context, data *money.UserBalancePayload) (*money.UserBalanceResponse, error) {
	userID, ok := common.ExtractUserIDFromHeader(ctx)
	if !ok {
		return &money.UserBalanceResponse{
			AccountNumber: userID,
			Balance:       "0",
		}, errors.New("invalid user ID")
	}
	amountStr, err := h.moneyService.PublicGetUserBalance(ctx, userID)
	return &money.UserBalanceResponse{
		AccountNumber: userID,
		Balance:       amountStr,
	}, err
}

func (h *Handler) GetUserBalancePrivate(ctx context.Context, data *money.UserBalancePayload) (*money.UserBalancePrivateResponse, error) {
	amount, _, err := h.moneyService.GetUserBalance(ctx, data.GetAccountNumber())
	return &money.UserBalancePrivateResponse{
		AccountNumber: data.GetAccountNumber(),
		Balance:       amount,
	}, err
}

func (h *Handler) UpdateUserBalance(ctx context.Context, data *money.UpdateBalancePayload) (*money.UpdateBalanceResponse, error) {
	err := h.moneyService.UpdateUserBalance(ctx, data.GetRequestId(), data.GetAccountNumber(), data.GetAmount())
	if err != nil {
		return &money.UpdateBalanceResponse{
			AccountNumber: data.GetAccountNumber(),
			Success:       false,
			Message:       err.Error(),
		}, err
	}

	finalAmount, _, err := h.moneyService.GetUserBalance(ctx, data.GetAccountNumber())
	var msg string
	if err != nil {
		msg = err.Error()
	}
	return &money.UpdateBalanceResponse{
		AccountNumber: data.GetAccountNumber(),
		Success:       err == nil,
		Balance:       finalAmount,
		Message:       msg,
	}, err
}
