package services

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
)

func (b BankingService) SubmitTransfer(ctx context.Context, payload model.BankingCashHistory) error {
	return nil
}

func (b BankingService) SubmitPayment(ctx context.Context, payload model.BankingCashHistory) error {
	return nil
}

func (b BankingService) SubmitDeposit(ctx context.Context, payload model.BankingCashHistory) error {
	return nil
}
