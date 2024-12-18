package services

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
)

func (s *SettlementService) HandleSettlementReport(ctx context.Context, accountNo, date string) error {
	return nil
}

func (s *SettlementService) executeSettlementReport(data *model.SettlementPayload) {
	// add trx into queue

}
