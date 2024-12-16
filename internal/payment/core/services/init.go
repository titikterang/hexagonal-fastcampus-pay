package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type PaymentService struct {
	config     *config.Config
	repository ports.PaymentRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.PaymentRepositoryAdapter) *PaymentService {
	return &PaymentService{
		config:     cfg,
		repository: repository,
	}
}
