package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"strconv"
	"time"
)

func (p *PaymentService) eligibleUser(ctx context.Context, accountNo string) bool {
	data, err := p.repository.GetAccountInfo(ctx, accountNo)
	if err != nil {
		log.Err(err).Msg("eligibleAccountStatus.repository.GetAccountInfo")
		return false
	}

	return data.Status == model.UserStatusActive
}

func (p *PaymentService) SubmitPaymentRequest(ctx context.Context, payload model.SubmitPaymentPayload) (string, error) {
	// cek user eligible info
	if !p.eligibleUser(ctx, payload.AccountNo) {
		return "", errors.New("user is not eligible to do payment")
	}
	// compose message to money service
	invID := fmt.Sprintf("INV/%d", time.Now().UnixNano())
	// insert into history
	p.repository.InsertPaymentHistory(ctx, model.PaymentInfo{
		InvoiceID:     invID,
		MerchantID:    payload.MerchantID,
		AccountNumber: payload.AccountNo,
		Amount:        decimal.NewFromInt(payload.Amount),
		DateTime:      time.Now(),
		Status:        model.Status_UNPAID,
	})

	// publish to money service
	msg := types.TransactionValidateInfo{
		TransactionID:       invID,
		Amount:              payload.Amount,
		SourceAccountNumber: payload.AccountNo,
		TransactionType:     strconv.Itoa(types.TransactionTypePayment),
		MerchantID:          payload.MerchantID,
	}

	err := p.repository.PublishPaymentValidateRequest(ctx, msg)

	return invID, err
}
func (p *PaymentService) GetPaymentInfoByID(ctx context.Context, id string) (model.PaymentDetails, error) {
	data, hist, err := p.repository.GetPaymentInfo(ctx, id)
	return model.PaymentDetails{
		Info:    data,
		History: hist,
	}, err
}
func (p *PaymentService) HandleTransactionValidationReply(ctx context.Context, data types.TransactionValidateReplyInfo) error {
	// handle idempotence
	if p.repository.EventIDExists(ctx, model.TypeValidateMoneyReply, data.SourceAccountNumber, data.TransactionID) {
		return nil
	}

	var (
		status = model.Status_PENDING
		msg    string
	)
	// set status to failed if insuf
	if !data.BalanceSufficient {
		status = model.Status_FAILED
		msg = "insufficient user balance"
	}

	// insert history
	p.repository.UpdatePaymentStatus(ctx, model.PaymentInfo{
		InvoiceID:     data.TransactionID,
		MerchantID:    data.MerchantID,
		AccountNumber: data.SourceAccountNumber,
		Message:       msg,
		Status:        status,
	})
	if !data.BalanceSufficient {
		return nil
	}

	// publish to bank service
	p.repository.PublishPaymentBankingRequest(ctx, types.PaymentBankExecution{
		TransactionID: data.TransactionID,
		MerchantID:    data.MerchantID,
		Amount:        0,
	})

	return nil
}
func (p *PaymentService) HandleBankCallbackReply(ctx context.Context, data types.PaymentBankReply) error {
	// handle idempotence
	if p.repository.EventIDExists(ctx, model.TypeValidateBankReply, "", data.TransactionID) {
		return nil
	}

	var (
		status = payment.PaymentStatus_FAILED
		info   = model.PaymentInfo{
			Message: data.Message,
		}
	)
	if data.Status == "ok" {
		status = payment.PaymentStatus_PAID
	}
	p.repository.InsertPaymentHistory(ctx, model.PaymentInfo{
		Status:  info.SetStatusFromProto(status),
		Message: data.Message,
	})

	return nil
}
