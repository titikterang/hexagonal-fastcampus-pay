package model

import (
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"time"
)

type PaymentStatus string

const (
	UserStatusActive  = "A"
	UserStatusSuspend = "S"

	// redis key
	DefaultIdempotenceTTL        = 1 * time.Hour
	IdempotencePaymentMoneyReply = "pym:idmp:money:%s:%s"
	IdempotencePaymentBankReply  = "pym:idmp:bank:%s:%s"

	MoneyTopicKey   = "money_service"
	BankingTopicKey = "banking_service"

	TypeValidateMoneyReply = "money"
	TypeValidateBankReply  = "bank"

	// status
	Status_UNPAID  PaymentStatus = "UNPAID"
	Status_PENDING PaymentStatus = "PENDING"
	Status_FAILED  PaymentStatus = "FAILED"
	Status_PAID    PaymentStatus = "PAID"
)

type SubmitPaymentPayload struct {
	MerchantID string
	Amount     int64
	AccountNo  string
}

type PaymentInfo struct {
	InvoiceID     string
	MerchantID    string
	AccountNumber string
	Message       string
	Amount        decimal.Decimal
	DateTime      time.Time
	Status        PaymentStatus
}

type PaymentHistory struct {
	DateTime string
	Status   PaymentStatus
}

type PaymentDetails struct {
	Info    PaymentInfo
	History []PaymentHistory
}

func (p PaymentInfo) SetStatusFromProto(status payment.PaymentStatus) PaymentStatus {
	switch status {
	case payment.PaymentStatus_UNPAID:
		return Status_UNPAID
	case payment.PaymentStatus_PENDING:
		return Status_PENDING
	case payment.PaymentStatus_FAILED:
		return Status_FAILED
	case payment.PaymentStatus_PAID:
		return Status_PAID
	}

	return Status_PENDING
}

func (p PaymentInfo) GetStatusProto() payment.PaymentStatus {
	switch p.Status {
	case Status_UNPAID:
		return payment.PaymentStatus_UNPAID
	case Status_PENDING:
		return payment.PaymentStatus_PENDING
	case Status_FAILED:
		return payment.PaymentStatus_FAILED
	case Status_PAID:
		return payment.PaymentStatus_PAID
	}

	return payment.PaymentStatus_UNSPECIFIED
}

type PaymentPrecheckInfo struct {
	Balance         int64
	UserProfileInfo types.UserProfileInfo
}
