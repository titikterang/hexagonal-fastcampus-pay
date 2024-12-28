package main

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/vault"
)

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	//FetchConfigFromVault(cfg)

	redisClient := InitRedis(cfg)
	dbClient, err := InitDB(cfg)
	if err != nil {
		return nil, err
	}

	repo := repository.NewDatastoreRepository(cfg, redisClient, dbClient)
	hdl, err := handler.NewHandler(cfg, services.NewService(cfg, repo))
	if err != nil {
		return nil, err
	}
	return hdl, nil
}

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}

func InitDB(cfg *config.Config) (postgre.DBInterface, error) {
	dbConn := postgre.InitDBConnection(cfg)
	client, err := dbConn.InitiateConnection()
	if err != nil {
		// log error
		log.Fatalf("failed to connect to db, err : %#v", err)
	}
	return client, nil
}

func FetchConfigFromVault(cfg *config.Config) {
	// init vault client
	client, err := vault.NewVaultClient(cfg.Vault.Token, cfg.Vault.Address)
	if err != nil {
		log.Fatalf("failed to connect to vault, err : %#v", err)
	}

	secret, err := client.GetSecretValue(context.Background(), cfg.Vault.Path)
	if err != nil {
		log.Fatalf("failed to GetSecretValue from vault, err : %#v", err)
	}

	var secretConfig config.SecretConfig
	data, err := json.Marshal(secret)
	if err != nil {
		log.Fatalf("failed to Marshal from vault, err : %#v", err)
	}

	err = json.Unmarshal(data, &secretConfig)
	if err != nil {
		log.Fatalf("failed to Unmarshal to secret config, err : %#v", err)
	}

	// overwrite from vault secret
	cfg.Token.KeyID = secretConfig.TokenConfig.KeyID
	cfg.Token.Secret = secretConfig.TokenConfig.Secret
	// overwrite db config
	cfg.Postgre.Username = secretConfig.Postgre.Username
	cfg.Postgre.Password = secretConfig.Postgre.Password
	cfg.Postgre.Dbname = secretConfig.Postgre.Dbname
}
