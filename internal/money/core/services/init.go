package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type MoneyService struct {
	config     *config.Config
	repository ports.MoneyRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.MoneyRepositoryAdapter) *MoneyService {
	return &MoneyService{
		config:     cfg,
		repository: repository,
	}
}
