package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
)

type MoneyService struct {
	config      *config.Config
	repository  ports.MoneyRepositoryAdapter
	currentTime types.CurrentTime
}

func NewService(cfg *config.Config, repository ports.MoneyRepositoryAdapter) *MoneyService {
	return &MoneyService{
		config:     cfg,
		repository: repository,
	}
}
