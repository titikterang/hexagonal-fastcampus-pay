package repository

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
)

type DatastoreRepository struct {
	cfg         *config.Config
	redisClient *redis.Client
	dbClient    postgre.DBInterface
	queries     statementQueries
}

func NewDatastoreRepository(cfg *config.Config, client *redis.Client, dbClient postgre.DBInterface) ports.DatastoreRepositoryAdapter {
	repo := &DatastoreRepository{
		cfg:         cfg,
		redisClient: client,
		dbClient:    dbClient,
	}
	err := repo.initDBSchema(context.TODO())
	if err != nil {
		log.Error(err)
	}

	err = repo.prepareStatements()
	if err != nil {
		log.Error(err)
	}
	return repo
}
