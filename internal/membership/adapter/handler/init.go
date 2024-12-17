package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/membership"
)

type Handler struct {
	membership.UnimplementedMembershipServiceServer
	config            *config.Config
	membershipService ports.MembershipServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.MembershipServiceAdapter) (*Handler, error) {
	return &Handler{
		config:            cfg,
		membershipService: adapter,
	}, nil
}
