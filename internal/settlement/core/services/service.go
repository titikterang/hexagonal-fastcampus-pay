package services

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

func (s *SettlementService) HandleSettlementReport(ctx context.Context, accountNo, date string) ([]model.SettlementPayload, error) {
	return nil, nil
}

func (s *SettlementService) executeSettlementReport(data *model.SettlementPayload) {
	ctx := context.Background()
	// idempotence check
	if s.repository.EventIDExists(ctx, data.TransactionID) {
		return
	}
	_ = s.repository.SaveEventID(ctx, data.TransactionID)

	// get trx info
	info, err := s.repository.GetTransactionInfoByID(ctx, data.TransactionID, data.SettlementType)
	if err != nil {
		log.Error().Msgf("failed to get trx info , %#v", err)
		return
	}

	// calculate settlement fee
	fee, feeAmount := s.getFee(data.SettlementType, info)
	// update to data store
	err = s.repository.SaveSettlementReport(ctx, model.SettlementPayload{
		AccountNo:      info.AccountNo,
		TransactionID:  info.ID,
		Amount:         data.Amount,
		FeeAmount:      feeAmount,
		FeePercentage:  fee,
		SettlementType: data.SettlementType,
	})
	if err != nil {
		log.Error().Msgf("failed to SaveSettlementReport , %#v", err)
		return
	}

	// update balance
	err = s.repository.UpdateUserBalance(ctx, info.AccountNo, feeAmount)
	if err != nil {
		log.Error().Msgf("failed to UpdateUserBalance , %#v", err)
		return
	}
}

func (s *SettlementService) getFee(sType model.SettlementType, info types.TransactionInfoResult) (fee decimal.Decimal, feeAmount decimal.Decimal) {
	fee = decimal.NewFromFloat(model.PaymentFee)
	if sType == model.TransferSettlement {
		fee = decimal.NewFromFloat(model.TransferFee)
	}

	feeAmount = fee.Div(decimal.NewFromFloat(100.0)).Mul(info.Amount)
	return fee, feeAmount
}
