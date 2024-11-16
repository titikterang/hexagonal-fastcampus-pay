package repository

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
)

func (r *BankingRepository) SaveHistoryToDB(ctx context.Context, data model.BankingCashHistory) error {
	return nil
}
