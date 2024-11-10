package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type MembershipService struct {
	config     *config.Config
	repository ports.DatastoreRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.DatastoreRepositoryAdapter) ports.MembershipServiceAdapter {
	return &MembershipService{
		config:     cfg,
		repository: repository,
	}
}
