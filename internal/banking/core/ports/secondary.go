package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type BankingRepositoryAdapter interface {
	// database repository
	SaveHistoryToDB(ctx context.Context, data model.BankingCashHistory) error

	// external API repository
	GetAccountInfo(ctx context.Context, accountNumber string) (types.UserProfileInfo, error)
}
