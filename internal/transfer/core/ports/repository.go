package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type TransferServiceRepositoryAdapter interface {
	// external api call
	GetAccountInfo(ctx context.Context, accountNumber string) (types.UserProfileInfo, error)

	// datastore
	SaveTransferHistory(ctx context.Context, data model.TransferInfo) error
	UpdateTransferHistory(ctx context.Context, status, ID string) error
	GetTransferHistory(ctx context.Context, filter string) ([]model.TransferInfo, error)

	// idempotence check
	EventIDExists(ctx context.Context, eventType model.EventType, accountNo, id string) bool
	SaveEventID(ctx context.Context, eventType model.EventType, accountNo, id string) error

	// publish to other service
	PublishTransferValidateRequest(ctx context.Context, info types.TransactionValidateInfo) error       // to money service
	PublishTransferBankingRequest(ctx context.Context, payload types.TransferExternalBankPayload) error // to banking service
}
