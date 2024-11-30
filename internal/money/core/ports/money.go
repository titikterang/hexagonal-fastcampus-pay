package ports

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"google.golang.org/genproto/googleapis/type/money"
)

type MoneyServiceAdapter interface {
	// return string
	PublicGetUserBalance(ctx context.Context, accountNumber string) (string, error)
	// return money
	GetUserBalance(ctx context.Context, accountNumber string) (*money.Money, decimal.Decimal, error)
	// payload money
	UpdateUserBalance(ctx context.Context, reqID, accountNumber string, amount *money.Money) error
}

type MoneyServiceConsumerAdapter interface {
	HandleTransactionValidation(ctx context.Context, reply types.TransactionValidateInfo) error
}
