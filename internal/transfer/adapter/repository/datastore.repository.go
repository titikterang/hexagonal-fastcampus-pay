package repository

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"time"
)

func (r *TransferRepository) SaveTransferHistory(ctx context.Context, data model.TransferInfo) error {
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

	_, err = r.queries.InsertTransferHistory.ExecContext(ctx, map[string]interface{}{
		"transaction_id":             data.TransactionId,
		"transfer_date":              time.Now().Format(time.DateOnly),
		"amount":                     data.Amount,
		"source_account_number":      data.SourceAccountNumber,
		"destination_account_number": data.DestinationAccountNumber,
		"transfer_type":              data.TransferType,
		"bank_code":                  data.BankCode,
		"status":                     data.Status,
		"created_at":                 time.Now(),
		"updated_at":                 time.Now(),
		"message":                    data.Message,
	})
	if err != nil {
		log.Error().Msgf("err save history %#v", err)
	}
	return err
}

func (r *TransferRepository) UpdateTransferHistory(ctx context.Context, status, ID string) error {
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

	_, err = r.queries.UpdateTransferStatus.ExecContext(ctx, map[string]interface{}{
		"status":  status,
		"updated": time.Now(),
		"id":      ID,
	})
	return err
}

func (r *TransferRepository) GetTransferHistory(ctx context.Context, filter string) ([]model.TransferInfo, error) {
	var (
		resultSet = []model.TransferInfoDB{}
		data      = []model.TransferInfo{}
	)
	err := r.queries.GetTransferHistory.SelectContext(ctx, &resultSet, map[string]interface{}{
		"transfer_date": filter,
	})

	for _, v := range resultSet {
		data = append(data, model.TransferInfo{
			TransactionId:            v.TransactionId,
			TransferDate:             v.TransferDate,
			Amount:                   v.Amount,
			SourceAccountNumber:      v.SourceAccountNumber,
			DestinationAccountNumber: v.DestinationAccountNumber,
			TransferType:             v.TransferType,
			BankCode:                 v.BankCode,
			Status:                   v.Status,
			CreatedAt:                v.CreatedAt,
			UpdatedAt:                v.UpdatedAt,
			Message:                  v.Message,
		})
	}
	return data, err
}

func (r *TransferRepository) EventIDExists(ctx context.Context, eventType model.EventType, accountNo, id string) bool {
	var key string
	switch eventType {
	case model.EventTypeSubmitTransfer:
		key = fmt.Sprintf(model.IdempotenceTransferHandler, accountNo, id)
	case model.EventTypeMoneyReply:
		key = fmt.Sprintf(model.IdempotenceMoneyReply, accountNo, id)
	case model.EventTypeBankReply:
		key = fmt.Sprintf(model.IdempotenceBankingReply, accountNo, id)
	default:
		return true
	}

	res, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	return res == "1"
}

func (r *TransferRepository) SaveEventID(ctx context.Context, eventType model.EventType, accountNo, id string) error {
	var key string
	switch eventType {
	case model.EventTypeSubmitTransfer:
		key = fmt.Sprintf(model.IdempotenceTransferHandler, accountNo, id)
	case model.EventTypeMoneyReply:
		key = fmt.Sprintf(model.IdempotenceMoneyReply, accountNo, id)
	case model.EventTypeBankReply:
		key = fmt.Sprintf(model.IdempotenceBankingReply, accountNo, id)
	default:
		return nil
	}

	return r.redisClient.SetEX(ctx, key, 1, model.DefaultIdempotenceTTL).Err()
}
