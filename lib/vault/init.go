package vault

import (
	"context"
	"github.com/hashicorp/vault-client-go"
	"github.com/rs/zerolog/log"
	"time"
)

type VaultClient struct {
	token  string
	host   string
	client *vault.Client
}

type VaultInterface interface {
	GetSecretValue(ctx context.Context, path string) (map[string]interface{}, error)
}

func NewVaultClient(token, host string) (VaultInterface, error) {
	client, err := vault.New(
		vault.WithAddress(host),
		vault.WithRequestTimeout(30*time.Second),
	)

	if err != nil {
		log.Fatal().Msgf("failed to connect to vault , %#v", err)
		return nil, err
	}

	if err := client.SetToken(token); err != nil {
		log.Fatal().Msgf("failed to init vault token, %#v", err)
		return nil, err
	}

	v := VaultClient{
		token:  token,
		host:   host,
		client: client,
	}

	return &v, nil
}
