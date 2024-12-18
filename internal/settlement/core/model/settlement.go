package model

import "github.com/shopspring/decimal"

type SettlementType string

const (
	PaymentFee  = 0.2
	TransferFee = 0.1

	TopicPayment  = "payment"
	TopicTransfer = "transfer"

	PaymentSettlement  SettlementType = "payment"
	TransferSettlement SettlementType = "transfer"
)

type SettlementPayload struct {
	AccountNo      string
	TransactionID  string
	Amount         decimal.Decimal
	SettlementType SettlementType
}
