package handler

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/transfer"
	"strconv"
)

func (h *Handler) GetTransferHistory(ctx context.Context, payload *transfer.TransferHistoryRequest) (*transfer.TransferHistoryResponse, error) {
	data, err := h.transferService.GetTransferHistory(ctx, payload.GetDate())
	if err != nil {

	}
	returnData := transfer.TransferHistoryResponse{
		TransferHistory: make([]*transfer.TransferBalanceResponse, 0),
	}
	for _, v := range data {
		returnData.TransferHistory = append(returnData.TransferHistory, &transfer.TransferBalanceResponse{
			Message: v.Message,
			Status:  v.Status,
			TrxId:   v.TransactionId,
			Amount:  v.Amount,
		})
	}
	return nil, nil
}

func (h *Handler) SubmitTransferBalance(ctx context.Context, payload *transfer.TransferBalanceRequest) (*transfer.TransferBalanceResponse, error) {
	ID, err := h.transferService.SubmitTransferBalance(ctx, model.TransferInfo{
		Amount:                   payload.GetAmount(),
		SourceAccountNumber:      payload.GetSourceAccountNumber(),
		DestinationAccountNumber: payload.GetDestinationAccount(),
		TransferType:             string(payload.GetTransferType()),
		BankCode:                 strconv.FormatInt(payload.GetAccountCode(), 10),
	})
	return &transfer.TransferBalanceResponse{
		Status: model.TransferStatusPending,
		TrxId:  ID,
		Amount: payload.GetAmount(),
	}, err
}
