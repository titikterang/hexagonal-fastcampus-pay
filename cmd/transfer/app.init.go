package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
)

func initHandler(cfg *config.Config) (*handler.Handler, *handler.ConsumerHandler, error) {
	masterClient, err := InitDBMaster(cfg)
	if err != nil {
		return nil, nil, err
	}
	slaveClient, err := InitDBSlave(cfg)
	if err != nil {
		return nil, nil, err
	}

	// init kafka client
	client, err := kafka.InitKafkaClient(cfg)
	if err != nil {
		if err != nil {
			log.Fatal("failed initiate NewHandler: %v", err)
		}
	}

	repo := repository.NewTransferRepository(cfg, masterClient, slaveClient, client)
	svc := services.NewService(cfg, repo)

	hdl, err := handler.NewHandler(cfg, svc)
	if err != nil {
		return nil, nil, err
	}

	consHandler, err := handler.NewConsumer(cfg, client, svc)
	if err != nil {
		return nil, nil, err
	}
	return hdl, consHandler, nil
}

func InitDBMaster(cfg *config.Config) (postgre.DBInterface, error) {
	dbConn := postgre.InitDBConnection(cfg)
	client, err := dbConn.InitiateMasterConnection()
	if err != nil {
		// log error
		log.Fatalf("failed to connect to db, err : %#v", err)
	}
	return client, nil
}
func InitDBSlave(cfg *config.Config) (postgre.DBInterface, error) {
	dbConn := postgre.InitDBConnection(cfg)
	client, err := dbConn.InitiateSlaveConnection()
	if err != nil {
		// log error
		log.Fatalf("failed to connect to db, err : %#v", err)
	}
	return client, nil
}

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}
