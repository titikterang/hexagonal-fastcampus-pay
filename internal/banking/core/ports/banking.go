package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type BankingServiceAdapter interface {
	SubmitTransfer(ctx context.Context, payload model.BankingCashHistory) error
	SubmitPayment(ctx context.Context, payload model.BankingCashHistory) error
	SubmitDeposit(ctx context.Context, payload model.BankingCashHistory) error
}

type BankingServiceConsumerAdapter interface {
	HandleBankPayment(ctx context.Context, data types.PaymentBankExecution) error
	HandleBankTransfer(ctx context.Context) error
}
