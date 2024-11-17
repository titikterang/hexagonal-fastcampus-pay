package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
)

func (r *BankingRepository) SaveHistoryToDB(ctx context.Context, data model.BankingCashHistory) error {
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

	_, err = r.queries.InsertCashHistory.ExecContext(ctx, map[string]interface{}{
		"account_source":      data.SourceAccount,
		"account_destination": data.DestinationAccount,
		"amount":              data.Amount,
		"cash_movemet_type":   data.CashMovementType,
	})
	return nil
}
