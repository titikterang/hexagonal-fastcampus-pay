package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type TransferServiceRepository interface {
	// external api call
	GetAccountInfo(ctx context.Context, accountNumber string) (types.UserProfileInfo, error)

	// datastore
	SaveTransferHistory()
	UpdateTransferHistory()

	// publish to other service
	PublishTransferValidateRequest() // to money service
	PublishTransferBankingRequest()  // to banking service
}
