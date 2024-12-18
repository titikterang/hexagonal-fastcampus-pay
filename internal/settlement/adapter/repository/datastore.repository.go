package repository

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// to datastore

func (r *SettlementRepository) SaveSettlementReport(ctx context.Context, data model.SettlementPayload) error {
	amount, err := bson.ParseDecimal128(data.Amount.String())
	if err != nil {
		return err
	}
	fee, err := bson.ParseDecimal128(data.FeePercentage.String())
	if err != nil {
		return err
	}
	feeAmount, err := bson.ParseDecimal128(data.FeeAmount.String())
	if err != nil {
		return err
	}
	dbPayload := model.SettlementReportDB{
		AccountNo:      data.AccountNo,
		TransactionID:  data.TransactionID,
		Amount:         amount,
		FeeAmount:      feeAmount,
		FeePercentage:  fee,
		SettlementType: data.SettlementType,
		SettlementDate: bson.NewDateTimeFromTime(time.Now()),
	}

	_, err = r.DB.Collection("settlement_report").InsertOne(ctx, dbPayload)
	return err
}

func (r *SettlementRepository) GetSettlementReport(ctx context.Context, accountNo, date string) ([]model.SettlementPayload, error) {
	var (
		history = make([]model.SettlementPayload, 0)
	)

	filter := bson.D{{
		"account_no", accountNo},
	}
	cursor, err := r.DB.Collection("settlement_report").Find(ctx, filter)
	if err != nil {
		return history, err
	}

	for cursor.Next(ctx) {
		var data model.SettlementReportDB
		err = cursor.Decode(&data)
		if err != nil {
			return history, err
		}

		amount, err := decimal.NewFromString(data.Amount.String())
		if err != nil {
			return history, err
		}
		fee, err := decimal.NewFromString(data.FeePercentage.String())
		if err != nil {
			return history, err
		}
		feeAmount, err := decimal.NewFromString(data.FeeAmount.String())
		if err != nil {
			return history, err
		}

		history = append(history, model.SettlementPayload{
			AccountNo:      data.AccountNo,
			TransactionID:  data.TransactionID,
			Amount:         amount,
			FeeAmount:      feeAmount,
			FeePercentage:  fee,
			SettlementType: data.SettlementType,
		})
	}

	return history, nil
}

func (r *SettlementRepository) EventIDExists(ctx context.Context, id string) bool {
	key := fmt.Sprintf("settle:%s", id)
	res, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	return res == "1"
}

func (r *SettlementRepository) SaveEventID(ctx context.Context, id string) error {
	key := fmt.Sprintf("settle:%s", id)
	return r.redisClient.SetEx(ctx, key, 1, types.DefaultIdempotenceTTL).Err()
}
