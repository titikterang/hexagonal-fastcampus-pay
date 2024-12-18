package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/adapter/repository"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/services"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
)

func initHandler(cfg *config.Config) (*handler.Handler, mongo.DBInterface, error) {
	dbCon, err := mongo.InitDBConnection(cfg)
	if err != nil {
		// log error
		log.Fatalf("failed to connect to mongo db, err : %#v", err)
	}

	repo := repository.NewSettlementRepository(cfg,
		nil,
		dbCon.DBClient,
		nil,
		nil,
		nil)
	svc := services.NewService(cfg, repo)

	hdl, err := handler.NewHandler(cfg, svc)
	if err != nil {
		return nil, nil, err
	}
	return hdl, dbCon.DBClient, nil
}
