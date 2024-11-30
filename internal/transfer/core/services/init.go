package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type TransferService struct {
	config     *config.Config
	repository ports.TransferServiceRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.TransferServiceRepositoryAdapter) *TransferService {
	return &TransferService{
		config:     cfg,
		repository: repository,
	}
}
