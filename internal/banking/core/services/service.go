package services

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

func (b BankingService) SubmitTransfer(ctx context.Context, payload model.BankingCashHistory) error {
	return b.repository.SaveHistoryToDB(ctx, payload)
}

func (b BankingService) SubmitPayment(ctx context.Context, payload model.BankingCashHistory) error {
	return b.repository.SaveHistoryToDB(ctx, payload)
}

func (b BankingService) SubmitDeposit(ctx context.Context, payload model.BankingCashHistory) error {
	payload.SourceAccount = "fastcampus_main_account"
	//  only submit deposit if user status == A
	data, err := b.repository.GetAccountInfo(ctx, payload.DestinationAccount)
	if err != nil {
		log.Err(err).Msg("SubmitDeposit.b.repository.GetAccountInfo")
		return err
	}

	// cek account status
	if data.Status != model.UserStatusActive {
		// insert history reject deposit
		payload.CashMovementType = model.MovementDepositReject
		_ = b.repository.SaveHistoryToDB(ctx, payload)
		return errors.New("non active user are prohibited to submit deposit")
	}

	return b.repository.SaveHistoryToDB(ctx, payload)
}

func (b BankingService) HandleBankPayment(ctx context.Context, data types.PaymentBankExecution) error {
	// simulate payment gateway callback
	err := b.repository.PublishPaymentCallbackReply(ctx, types.PaymentBankReply{
		TransactionID: data.TransactionID,
		Status:        "ok",
		Message:       "Payment success",
	})
	return err
}
func (b BankingService) HandleBankTransfer(ctx context.Context) error {
	return nil
}
