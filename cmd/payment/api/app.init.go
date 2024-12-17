package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/membership"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
)

func initHandler(cfg *config.Config) (*handler.Handler, kafka.KafkaClientInterface, mongo.DBInterface, error) {
	dbCon, err := mongo.InitDBConnection(cfg)
	if err != nil {
		// log error
		log.Fatalf("failed to connect to mongo db, err : %#v", err)
	}

	// init kafka client producer
	clientProducer, err := kafka.InitKafkaClient(cfg, kafka.TypeProducerClient)
	if err != nil {
		if err != nil {
			log.Fatal("failed initiate clientProducer: %v", err)
		}
	}

	// init grpc client
	connMembership, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(cfg.ExternalAPI.MembershipGrpc),
	)

	connMoney, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(cfg.ExternalAPI.MoneyGrpc),
	)

	membershipClient := membership.NewMembershipServiceClient(connMembership)
	moneyClient := money.NewMoneyServiceClient(connMoney)

	redisClient := InitRedis(cfg)
	repo := repository.NewPaymentRepository(cfg,
		redisClient,
		dbCon.DBClient,
		clientProducer,
		membershipClient,
		moneyClient)
	svc := services.NewService(cfg, repo)

	hdl, err := handler.NewHandler(cfg, svc)
	if err != nil {
		return nil, nil, nil, err
	}
	return hdl, clientProducer, dbCon.DBClient, nil
}

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}
