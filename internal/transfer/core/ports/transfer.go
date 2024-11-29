package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type TransferServiceAPIAdapter interface {
	SubmitTransferBalance(ctx context.Context, data model.TransferInfo) (string, error)
	GetTransferHistory(ctx context.Context, filter string) ([]model.TransferInfo, error)
}

type TransferServiceConsumerAdapter interface {
	HandleTransactionValidationReply(ctx context.Context, reply types.TransactionValidateReplyInfo) error
	HandleBankCallbackReply(ctx context.Context, reply types.TransferExternalBankReply) error
}
