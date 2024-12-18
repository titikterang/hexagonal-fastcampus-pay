package types

import "github.com/shopspring/decimal"

type TransactionInfoResult struct {
	AccountNo string
	Amount    decimal.Decimal
	ID        string
}
