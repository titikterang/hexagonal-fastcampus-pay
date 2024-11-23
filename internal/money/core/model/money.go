package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type CashMovementInfo struct {
	RequestID     string
	AccountNumber string
	Date          time.Time
	Amount        decimal.Decimal
}
