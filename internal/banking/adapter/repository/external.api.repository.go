package repository

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
)

func (r *BankingRepository) GetAccountInfo(ctx context.Context, accountNumber string) (model.UserProfileInfo, error) {
	return model.UserProfileInfo{}, nil
}
