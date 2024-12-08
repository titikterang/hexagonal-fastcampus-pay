package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
)

func initHandler(cfg *config.Config) (*handler.ConsumerHandler, kafka.KafkaClientInterface, error) {
	masterClient, err := InitDBMaster(cfg)
	if err != nil {
		return nil, nil, err
	}
	slaveClient, err := InitDBSlave(cfg)
	if err != nil {
		return nil, nil, err
	}

	redisClient := InitRedis(cfg)

	// init kafka client consumer
	clientConsumer, err := kafka.InitKafkaClient(cfg, kafka.TypeConsumerClient)
	if err != nil {
		if err != nil {
			log.Fatal("failed initiate clientConsumer : %v", err)
		}
	}

	// init kafka client producer
	clientProducer, err := kafka.InitKafkaClient(cfg, kafka.TypeProducerClient)
	if err != nil {
		if err != nil {
			log.Fatal("failed initiate clientProducer: %v", err)
		}
	}

	repo := repository.NewTransferRepository(cfg, redisClient, masterClient, slaveClient, clientProducer)
	svc := services.NewService(cfg, repo)

	consHandler, err := handler.NewConsumer(cfg, clientConsumer, svc)
	if err != nil {
		return nil, nil, err
	}
	return consHandler, clientProducer, nil
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
