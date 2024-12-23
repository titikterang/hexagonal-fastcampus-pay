package ports

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
)

type SettlementServiceAdapter interface {
	HandleSettlementReport(ctx context.Context, accountNo, date string) ([]model.SettlementPayload, error)
}

type SettlementConsumerAdapter interface {
	HandleSettlementQueue(data model.SettlementPayload)
	InitWorkerPool() // only initiate worker pool on settlement consumer
}
