package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
)

func initHandler(cfg *config.Config) (*handler.Handler, error) {
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
