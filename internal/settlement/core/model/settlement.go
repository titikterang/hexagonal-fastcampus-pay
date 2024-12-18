package model

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

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
	FeeAmount      decimal.Decimal
	FeePercentage  decimal.Decimal
	SettlementType SettlementType
}

type SettlementReportDB struct {
	AccountNo      string          `bson:"account_no"`
	TransactionID  string          `bson:"transaction_id"`
	Amount         bson.Decimal128 `bson:"amount"`
	FeeAmount      bson.Decimal128 `bson:"fee_amount"`
	FeePercentage  bson.Decimal128 `bson:"fee_percentage"`
	SettlementType SettlementType  `bson:"settlement_type"`
	SettlementDate bson.DateTime   `bson:"settlement_date"`
}
