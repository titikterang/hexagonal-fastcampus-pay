package repository

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/common"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/transfer"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"strconv"
	"time"
)

func (r *SettlementRepository) UpdateUserBalance(ctx context.Context, accountNo string, amount decimal.Decimal) error {
	_, err := r.moneyClient.UpdateUserBalance(ctx, &money.UpdateBalancePayload{
		AccountNumber: accountNo,
		Amount:        common.DecimalToMoney(amount.Mul(decimal.NewFromInt(-1))),
		RequestId:     strconv.FormatInt(time.Now().UnixNano(), 10),
	})
	return err
}

func (r *SettlementRepository) GetTransactionInfoByID(ctx context.Context, id string, settleType model.SettlementType) (types.TransactionInfoResult, error) {
	var (
		result types.TransactionInfoResult
		err    error
	)

	if settleType == model.PaymentSettlement {
		resp, err := r.paymentClient.GetPaymentInfoByID(ctx, &payment.PaymentInfoPayload{
			TrxId: id,
		})
		if err != nil {
			return result, err
		}

		return types.TransactionInfoResult{
			AccountNo: resp.AccountNo,
			Amount:    common.MoneyToDecimal(resp.Amount),
			ID:        resp.TrxId,
		}, nil
	}

	resp, err := r.transferClient.GetTransferInfoByID(ctx, &transfer.TransferInfoPayload{
		TrxId: id,
	})
	if err != nil {
		return result, err
	}

	return types.TransactionInfoResult{
		AccountNo: resp.AccountNo,
		Amount:    common.MoneyToDecimal(resp.Amount),
		ID:        resp.TrxId,
	}, err
}
