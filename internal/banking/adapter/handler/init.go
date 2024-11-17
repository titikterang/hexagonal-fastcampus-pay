package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type Handler struct {
	config         *config.Config
	bankingService ports.BankingServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.BankingServiceAdapter) (*Handler, error) {
	return &Handler{
		config:         cfg,
		bankingService: adapter,
	}, nil
}
