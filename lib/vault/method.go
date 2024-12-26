package vault

import (
	"context"
	"github.com/hashicorp/vault-client-go"
	"github.com/rs/zerolog/log"
)

func (v *VaultClient) GetSecretValue(ctx context.Context, path string) (map[string]interface{}, error) {
	secret, err := v.client.Secrets.KvV2Read(ctx, path, vault.WithMountPath("secret"))
	if err != nil {
		log.Error().Msgf("failed to fetch secret data, %#v", err)
	}
	return secret.Data.Data, nil
}
