package ports

import (
	"context"
	"google.golang.org/genproto/googleapis/type/money"
)

type MoneyServiceAdapter interface {
	// return string
	PublicGetUserBalance(ctx context.Context, accountNumber string) (string, error)
	// return money
	GetUserBalance(ctx context.Context, accountNumber string) (*money.Money, error)
	// payload money
	UpdateUserBalance(ctx context.Context, reqID, accountNumber string, amount *money.Money) error
}
