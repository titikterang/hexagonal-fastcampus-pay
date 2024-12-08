package model

import (
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"time"
)

const (
	MoneyTopicKey   = "money_service"
	BankingTopicKey = "banking_service"
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
	Status        payment.PaymentStatus
}

type PaymentHistory struct {
	DateTime string
	Status   payment.PaymentStatus
}

type PaymentDetails struct {
	Info    PaymentInfo
	History []PaymentHistory
}
