package services

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/common"
	"google.golang.org/genproto/googleapis/type/money"
	"time"
)

// PublicGetUserBalance - only get from redis
// deps, redis get
func (s *MoneyService) PublicGetUserBalance(ctx context.Context, accountNumber string) (string, error) {
	// if empty, then recalculate
	res, err := s.repository.GetSnapshot(ctx, accountNumber)
	if err != nil {
		s.updateSnapshoot(ctx, accountNumber)
		return s.repository.GetSnapshot(ctx, accountNumber)
	}

	return res, err
}

func (s *MoneyService) updateSnapshoot(ctx context.Context, accountNumber string) {
	finalAmount, err := s.GetUserBalance(ctx, accountNumber)
	if err != nil {
		log.Error().Msgf("failed to get user balance, %#v", err)
		return
	}
	_ = s.repository.UpdateSnapshot(ctx, accountNumber, common.MoneyToString(finalAmount))
}

// GetUserBalance - get from redis and re calculate
// deps redis hgetall
func (s *MoneyService) GetUserBalance(ctx context.Context, accountNumber string) (*money.Money, error) {
	data, err := s.repository.GetCashMovementFromCache(ctx, accountNumber)
	if err != nil {
		return nil, err
	}

	var amount decimal.Decimal
	for _, v := range data {
		amount = amount.Add(v.Amount)
	}

	return common.DecimalToMoney(amount), nil
}

// append into db cash movement, append to redis cash movement, update snapshoot
// insert db
// hset redis
// set snapshot
func (s *MoneyService) UpdateUserBalance(ctx context.Context, requestID, accountNumber string, amount *money.Money) error {
	if amount == nil {
		return errors.New("empty amount, skip update")
	}
	// validate reqID
	if s.repository.ReqIDExists(ctx, accountNumber, requestID) {
		return errors.New("reqID duplicate, skip update")
	}

	s.repository.SaveReqID(ctx, accountNumber, requestID)
	// update to datastore db & redis
	amountStr := common.MoneyToString(amount)
	amountInfo, err := decimal.NewFromString(amountStr)
	if err != nil {
		return errors.New("failed to get amount, skip update")
	}
	info := model.CashMovementInfo{
		RequestID:     requestID,
		AccountNumber: accountNumber,
		Date:          time.Now(),
		Amount:        amountInfo,
	}
	err = s.repository.AppendCashMovementIntoDatastore(ctx, info)
	if err != nil {
		return errors.New("failed to update into datastore, skip update")
	}

	// update to redis
	err = s.repository.AppendCashMovementInfoIntoCache(ctx, info)
	if err != nil {
		return errors.New("failed to update into cache, skip update")
	}

	// update snapshoot
	s.updateSnapshoot(ctx, accountNumber)
	return nil
}
