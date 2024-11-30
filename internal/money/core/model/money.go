package model

import (
	"github.com/shopspring/decimal"
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
}
