package services

import (
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"os"
)

type MembershipService struct {
	config     *config.Config
	repository ports.DatastoreRepositoryAdapter
	privKey    []byte
	pubKey     []byte
}

func NewService(cfg *config.Config, repository ports.DatastoreRepositoryAdapter) ports.MembershipServiceAdapter {
	priv, err := os.ReadFile(cfg.Token.PrivateKey)
	if err != nil {
		log.Error().Msgf("Failed to load secret file , err %#v", err)
	}
	pub, err := os.ReadFile(cfg.Token.PublicKey)
	if err != nil {
		log.Error().Msgf("Failed to load secret file , err %#v", err)
	}
	return &MembershipService{
		config:     cfg,
		repository: repository,
		privKey:    priv,
		pubKey:     pub,
	}
}
