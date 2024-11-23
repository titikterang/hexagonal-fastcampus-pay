package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
)

type MoneyRepository struct {
	cfg            *config.Config
	dbClientMaster postgre.DBInterface
	dbClientSlave  postgre.DBInterface
	queries        statementQueries
	redisClient    *redis.Client
}

func NewMoneyRepository(cfg *config.Config, redisClient *redis.Client, masterClient, slaveClient postgre.DBInterface) ports.MoneyDataStoreAdapter {
	repo := &MoneyRepository{
		cfg:            cfg,
		dbClientMaster: masterClient,
		dbClientSlave:  slaveClient,
		queries:        statementQueries{},
		redisClient:    redisClient,
	}
	//err := repo.initDBSchema(context.TODO())
	//if err != nil {
	//	log.Err(err)
	//}
	//
	//err = repo.prepareStatements()
	//if err != nil {
	//	log.Err(err)
	//}
	return repo
}
