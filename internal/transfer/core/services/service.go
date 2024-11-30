package services

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"strconv"
	"time"
)

func (s *TransferService) SubmitTransferBalance(ctx context.Context, data model.TransferInfo) (string, error) {
	// idempotence check

	// validate source account status

	// validate destination account , if transfer between fastcampus

	data.TransactionId = strconv.FormatInt(time.Now().UnixNano(), 10)
	err := s.repository.SaveTransferHistory(ctx, data)
	s.repository.PublishTransferValidateRequest(ctx, types.TransactionValidateInfo{
		TransactionID:       data.TransactionId,
		Amount:              data.Amount,
		SourceAccountNumber: data.SourceAccountNumber,
		TransactionType:     1,
		Destination:         data.DestinationAccountNumber,
		BankCode:            data.BankCode,
	})

	return data.TransactionId, err
}

func (s *TransferService) GetTransferHistory(ctx context.Context, filter string) ([]model.TransferInfo, error) {
	return s.repository.GetTransferHistory(ctx, filter)
}

func (s *TransferService) HandleTransactionValidationReply(ctx context.Context, reply types.TransactionValidateReplyInfo) error {
	// check validation result
	var finalStatus = model.TransferStatusRejected
	if reply.BalanceSufficient {
		finalStatus = model.TransferStatusPendingBank
	}

	if reply.BankCode == model.TransferIntraAccount {
		finalStatus = model.TransferStatusSuccess
	}

	// update history
	err := s.repository.UpdateTransferHistory(ctx, finalStatus, reply.ReplyID)
	if err != nil {
		log.Error().Msgf("error UpdateTransferHistory %#v", err)
	}

	// produce to bank if sufficient
	if !reply.BalanceSufficient ||
		(reply.BalanceSufficient && reply.BankCode == model.TransferIntraAccount) {
		return nil
	}

	// only produce to bank consumer if only transfer to other bank
	err = s.repository.PublishTransferBankingRequest(ctx, types.TransferExternalBankPayload{
		TransactionInfo:          reply.TransactionValidateInfo,
		DestinationAccountNumber: reply.Destination,
		BankCode:                 reply.BankCode,
	})
	return err
}

func (s *TransferService) HandleBankCallbackReply(ctx context.Context, reply types.TransferExternalBankReply) error {
	// update history
	err := s.repository.UpdateTransferHistory(ctx, reply.Status, reply.TransactionID)
	if err != nil {
		log.Error().Msgf("error UpdateTransferHistory %#v", err)
	}

	return nil
}
