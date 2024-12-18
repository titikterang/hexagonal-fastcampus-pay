package repository

import (
	"context"
	"github.com/shopspring/decimal"
)

func (r *SettlementRepository) UpdateUserBalance(ctx context.Context, accountNo string, amount decimal.Decimal) error {
	return nil
}
