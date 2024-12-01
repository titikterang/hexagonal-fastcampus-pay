package repository

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/model"
	"time"
)

// method to redis cache
func (r *MoneyRepository) GetSnapshot(ctx context.Context, accountNumber string) (string, error) {
	key := fmt.Sprintf("money:snapshoot:%s", accountNumber)
	return r.redisClient.Get(ctx, key).Result()
}

func (r *MoneyRepository) UpdateSnapshot(ctx context.Context, accountNumber, amount string) error {
	key := fmt.Sprintf("money:snapshoot:%s", accountNumber)
	return r.redisClient.Set(ctx, key, amount, 1*time.Hour).Err()
}

func (r *MoneyRepository) GetCashMovementFromCache(ctx context.Context, accountNumber string) ([]model.CashMovementInfo, error) {
	key := fmt.Sprintf("money:movement:%s", accountNumber)
	data, err := r.redisClient.HGetAll(context.TODO(), key).Result()
	if err != nil {
		return []model.CashMovementInfo{}, err
	}
	result := []model.CashMovementInfo{}
	for k, v := range data {
		amount, err := decimal.NewFromString(v)
		if err != nil {
			continue
		}
		result = append(result, model.CashMovementInfo{
			RequestID:     k,
			AccountNumber: accountNumber,
			Amount:        amount,
		})
	}
	return result, nil
}

// hset
func (r *MoneyRepository) AppendCashMovementInfoIntoCache(ctx context.Context, info model.CashMovementInfo) error {
	key := fmt.Sprintf("money:movement:%s", info.AccountNumber)
	return r.redisClient.HSet(ctx, key, map[string]interface{}{
		info.RequestID: info.Amount.InexactFloat64(),
	}).Err()
}

func (r *MoneyRepository) ReqIDExists(ctx context.Context, accountNumber, reqID string) bool {
	key := fmt.Sprintf("money:reqid:%s:%s", accountNumber, reqID)
	res, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	return res == "1"
}

// set to redis TTL 1jam
func (r *MoneyRepository) SaveReqID(ctx context.Context, accountNumber, reqID string) {
	key := fmt.Sprintf("money:reqid:%s:%s", accountNumber, reqID)
	r.redisClient.SetEx(ctx, key, 1, 1*time.Hour)
	return
}

// method to db
func (r *MoneyRepository) AppendCashMovementIntoDatastore(ctx context.Context, info model.CashMovementInfo) error {
	trx, err := r.dbClientMaster.Beginx()
	if err != nil {
		log.Err(err)
		return err
	}

	defer func() {
		if err != nil {
			err = trx.Rollback()
		} else {
			err = trx.Commit()
		}
	}()

	_, err = r.queries.InsertCashMovement.ExecContext(ctx, map[string]interface{}{
		"request_id":         info.RequestID,
		"account_number":     info.AccountNumber,
		"cash_movement_date": info.Date.Format(time.DateOnly),
		"amount":             info.Amount.InexactFloat64(),
		"cash_movement_type": "deposit",
	})
	return nil
}
