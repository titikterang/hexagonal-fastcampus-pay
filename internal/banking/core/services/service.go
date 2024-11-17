package services

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
)

func (b BankingService) SubmitTransfer(ctx context.Context, payload model.BankingCashHistory) error {
	return b.repository.SaveHistoryToDB(ctx, payload)
}

func (b BankingService) SubmitPayment(ctx context.Context, payload model.BankingCashHistory) error {
	return b.repository.SaveHistoryToDB(ctx, payload)
}

func (b BankingService) SubmitDeposit(ctx context.Context, payload model.BankingCashHistory) error {
	payload.SourceAccount = "fastcampus_main_account"
	return b.repository.SaveHistoryToDB(ctx, payload)
}
