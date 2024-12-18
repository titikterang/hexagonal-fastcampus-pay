package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
)

func initHandler(cfg *config.Config) (*handler.Handler, mongo.DBInterface, error) {
	masterClient, err := InitDBMaster(cfg)
	if err != nil {
		return nil, nil, err
	}
	slaveClient, err := InitDBSlave(cfg)
	if err != nil {
		return nil, nil, err
	}

	dbCon, err := mongo.InitDBConnection(cfg)
	if err != nil {
		// log error
		log.Fatalf("failed to connect to mongo db, err : %#v", err)
	}

	redisClient := InitRedis(cfg)

	repo := repository.NewMoneyRepository(cfg, redisClient, dbCon.DBClient, masterClient, slaveClient, nil)
	svc := services.NewService(cfg, repo)
	hdl, err := handler.NewHandler(cfg, svc)
	if err != nil {
		return nil, nil, err
	}
	return hdl, dbCon.DBClient, nil
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
