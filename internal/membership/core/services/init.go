package services

import (
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"os"
)

type keyPair struct {
	privKey []byte
	pubKey  []byte
}

type MembershipService struct {
	config         *config.Config
	repository     ports.DatastoreRepositoryAdapter
	authKeyPair    keyPair
	refreshKeypair keyPair
}

func NewService(cfg *config.Config, repository ports.DatastoreRepositoryAdapter) ports.MembershipServiceAdapter {
	// auth keypair
	priv, err := os.ReadFile(cfg.Token.PrivateKey)
	if err != nil {
		log.Error().Msgf("Failed to load PrivateKey file , err %#v", err)
	}
	pub, err := os.ReadFile(cfg.Token.PublicKey)
	if err != nil {
		log.Error().Msgf("Failed to load PublicKey file , err %#v", err)
	}
	// refresh keypair
	refreshPriv, err := os.ReadFile(cfg.Token.RefreshPrivateKey)
	if err != nil {
		log.Error().Msgf("Failed to load RefreshPrivateKey file , err %#v", err)
	}
	refreshPub, err := os.ReadFile(cfg.Token.RefreshPublicKey)
	if err != nil {
		log.Error().Msgf("Failed to load RefreshPublicKey file , err %#v", err)
	}

	return &MembershipService{
		config:     cfg,
		repository: repository,
		authKeyPair: keyPair{
			privKey: priv,
			pubKey:  pub,
		},
		refreshKeypair: keyPair{
			privKey: refreshPriv,
			pubKey:  refreshPub,
		},
	}
}
