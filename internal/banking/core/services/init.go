package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type BankingService struct {
	config     *config.Config
	repository ports.BankingRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.BankingRepositoryAdapter) ports.BankingServiceAdapter {
	return &BankingService{
		config:     cfg,
		repository: repository,
	}
}
