package model

import "time"

const (
	MovementDeposit       = "deposit"
	MovementDepositReject = "deposit_reject"
	MovementPayment       = "payment"
	MovementTransfer      = "transfer"
)

type BankingCashHistory struct {
	ID                 int       `json:"id" db:"id"`
	SourceAccount      string    `json:"source_account" db:"source_account"`
	DestinationAccount string    `json:"destination_account" db:"destination_account"`
	Amount             float64   `json:"amount" db:"amount"`
	CashMovementType   string    `json:"cash_movement_type" db:"cash_movement_type"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
}
