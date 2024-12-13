package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
)

func initHandler(cfg *config.Config) (*handler.ConsumerHandler, kafka.KafkaClientInterface, mongo.DBInterface, error) {
	dbCon, err := mongo.InitDBConnection(cfg)
	if err != nil {
		// log error
		log.Fatalf("failed to connect to mongo db, err : %#v", err)
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

	repo := repository.NewPaymentRepository(cfg, redisClient, dbCon.DBClient, clientProducer)
	svc := services.NewService(cfg, repo)

	consHandler, err := handler.NewConsumer(cfg, clientConsumer, svc)
	if err != nil {
		return nil, nil, nil, err
	}
	return consHandler, clientProducer, dbCon.DBClient, nil
}

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}
