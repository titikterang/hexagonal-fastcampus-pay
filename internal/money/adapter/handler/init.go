package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type Handler struct {
	config       *config.Config
	moneyService ports.MoneyServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.MoneyServiceAdapter) (*Handler, error) {
	return &Handler{
		config:       cfg,
		moneyService: adapter,
	}, nil
}
