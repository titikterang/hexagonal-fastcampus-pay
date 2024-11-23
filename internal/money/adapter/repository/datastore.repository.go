package repository

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/model"
)

// method to redis cache
func (r *MoneyRepository) GetSnapshot(ctx context.Context, accountNumber string) (string, error) {
	return "", nil
}
func (r *MoneyRepository) UpdateSnapshot(ctx context.Context, accountNumber, amount string) error {
	return nil
}
func (r *MoneyRepository) GetCashMovementFromCache(ctx context.Context, accountNumber string) ([]model.CashMovementInfo, error) {
	return nil, nil
}
func (r *MoneyRepository) AppendCashMovementInfoIntoCache(ctx context.Context, info model.CashMovementInfo) error {
	return nil
}

func (r *MoneyRepository) ReqIDExists(ctx context.Context, accountNumber, reqID string) bool {
	return false
}

func (r *MoneyRepository) SaveReqID(ctx context.Context, accountNumber, reqID string) {
	return
}

// method to db
func (r *MoneyRepository) AppendCashMovementIntoDatastore(ctx context.Context, info model.CashMovementInfo) error {

	return nil
}
