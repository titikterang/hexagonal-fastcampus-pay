package model

import "time"

const (
	TransferIntraAccount      = "1"
	TransferStatusPending     = "transfer_pending"
	TransferStatusPendingBank = "pending_bank_confirm"
	TransferStatusRejected    = "transfer_rejected"
	TransferStatusSuccess     = "transfer_success"

	TransferType_UNSPECIFIED     = "UNSPECIFIED"
	TransferType_BETWEEN_ACCOUNT = "BETWEEN_ACCOUNT"
	TransferType_BETWEEN_BANK    = "BETWEEN_BANK"
)

type TransferInfoDB struct {
	TransactionId            string    `json:"transaction_id" db:"transaction_id"`
	TransferDate             string    `json:"transfer_date" db:"transfer_date"`
	Amount                   int64     `json:"amount" db:"amount"`
	SourceAccountNumber      string    `json:"source_account_number" db:"source_account_number"`
	DestinationAccountNumber string    `json:"destination_account_number" db:"destination_account_number"`
	TransferType             string    `json:"transfer_type" db:"transfer_type"`
	BankCode                 string    `json:"bank_code" db:"bank_code"`
	Status                   string    `json:"status" db:"status"`
	CreatedAt                time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                time.Time `json:"updated_at" db:"updated_at"`
	Message                  string    `json:"message" db:"message"`
}

type TransferInfo struct {
	RequestID                string    `json:"request_id"`
	TransactionId            string    `json:"transaction_id"`
	TransferDate             string    `json:"transfer_date"`
	Amount                   int64     `json:"amount"`
	SourceAccountNumber      string    `json:"source_account_number"`
	DestinationAccountNumber string    `json:"destination_account_number"`
	TransferType             string    `json:"transfer_type"`
	BankCode                 string    `json:"bank_code"`
	Status                   string    `json:"status"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	Message                  string    `json:"message"`
}
