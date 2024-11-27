package ports

import "context"

type TransferServiceAdapter interface {
	// handler rest api
	SubmitTransferBalance(ctx context.Context) error
	GetTransferHistory(ctx context.Context) error

	// handler consumer
	HandleTransactionValidationReply(ctx context.Context) error
	HandleBankCallbackReply(ctx context.Context) error
}
