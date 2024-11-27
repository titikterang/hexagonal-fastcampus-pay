package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
)

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	masterClient, err := InitDBMaster(cfg)
	if err != nil {
		return nil, err
	}
	slaveClient, err := InitDBSlave(cfg)
	if err != nil {
		return nil, err
	}

	repo := repository.NewBankingRepository(cfg, masterClient, slaveClient)
	hdl, err := handler.NewHandler(cfg, services.NewService(cfg, repo))
	if err != nil {
		return nil, err
	}

	return hdl, nil
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
