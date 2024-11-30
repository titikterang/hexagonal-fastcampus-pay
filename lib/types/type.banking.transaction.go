package types

type TransferExternalBankPayload struct {
	TransactionInfo          TransactionValidateInfo `json:"transaction_info"`
	DestinationAccountNumber string                  `json:"destination_account_number"`
	BankCode                 string                  `json:"bank_code"`
}

type TransferExternalBankReply struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}
