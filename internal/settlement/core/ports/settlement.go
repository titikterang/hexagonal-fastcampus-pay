package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
)

type SettlementServiceAdapter interface {
	HandleSettlementReport(ctx context.Context, accountNo, date string) error
}

type SettlementConsumerAdapter interface {
	HandleSettlementQueue(data model.SettlementPayload)
}
