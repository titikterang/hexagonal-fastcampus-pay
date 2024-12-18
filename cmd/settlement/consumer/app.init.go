package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/transfer"
)

func initHandler(cfg *config.Config) (*handler.ConsumerHandler, mongo.DBInterface, error) {
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

	// grpc client
	connMoney, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(cfg.ExternalAPI.MoneyGrpc),
	)
	moneyClient := money.NewMoneyServiceClient(connMoney)

	connTransfer, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(cfg.ExternalAPI.TransferGrpc),
	)
	transferClient := transfer.NewTransferServiceClient(connTransfer)

	connPayment, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(cfg.ExternalAPI.PaymentGrpc),
	)
	paymentClient := payment.NewPaymentServiceClient(connPayment)

	repo := repository.NewSettlementRepository(cfg,
		redisClient,
		dbCon.DBClient,
		moneyClient,
		transferClient,
		paymentClient)
	svc := services.NewService(cfg, repo)

	consHandler, err := handler.NewConsumer(cfg, clientConsumer, svc)
	if err != nil {
		return nil, nil, err
	}
	return consHandler, dbCon.DBClient, nil
}

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}
