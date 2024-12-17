package model

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

const (
	TopicTrxValidate = "transaction_validate"
	TopicMoneyDLQ    = "dead_letter_queue"
)

type CashMovementInfo struct {
	RequestID     string
	AccountNumber string
	Date          time.Time
	Amount        decimal.Decimal
	MovementType  string
}

type UserCashInfo struct {
	AccountNumber string
	LastUpdate    time.Time
	BalanceAmount decimal.Decimal
}

type UserCashDBInfo struct {
	AccountNumber string          `bson:"account_no"`
	LastUpdate    time.Time       `bson:"last_update"`
	BalanceAmount bson.Decimal128 `bson:"balance_amount"`
}
