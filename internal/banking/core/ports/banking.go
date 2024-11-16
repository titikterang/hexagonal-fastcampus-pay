package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
)

type BankingServiceAdapter interface {
	SubmitTransfer(ctx context.Context, payload model.BankingCashHistory) error
	SubmitPayment(ctx context.Context, payload model.BankingCashHistory) error
	SubmitDeposit(ctx context.Context, payload model.BankingCashHistory) error
}
