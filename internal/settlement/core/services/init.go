package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type SettlementService struct {
	settlementQueue chan *model.SettlementPayload
	config          *config.Config
	repository      ports.SettlementRepository
}

func NewService(cfg *config.Config, repository ports.SettlementRepository) *SettlementService {
	s := SettlementService{
		config:          cfg,
		repository:      repository,
		settlementQueue: make(chan *model.SettlementPayload, cfg.WorkerPool.Limit),
	}
	s.initWorkerPool()
	return &s
}
