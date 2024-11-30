package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type MoneyService struct {
	config     *config.Config
	repository ports.MoneyDataStoreAdapter
}

func NewService(cfg *config.Config, repository ports.MoneyDataStoreAdapter) ports.MoneyServiceAdapter {
	return &MoneyService{
		config:     cfg,
		repository: repository,
	}
}
