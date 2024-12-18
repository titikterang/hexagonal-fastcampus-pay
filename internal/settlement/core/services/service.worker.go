package services

import (
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
)

func (s *SettlementService) HandleSettlementQueue(data model.SettlementPayload) {
	s.settlementQueue <- &data
}

func (s *SettlementService) initWorkerPool() {
	for w := 0; w < s.config.WorkerPool.Limit; w++ {
		s.runWorkerPool(w)
	}
}

func (s *SettlementService) runWorkerPool(idx int) {
	go func() {
		log.Info().Msgf("run worker .. %d", idx)
		for {
			select {
			case payload := <-s.settlementQueue:
				s.executeSettlementReport(payload)
			}
		}
	}()
}
