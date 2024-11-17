package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
	"net/http"
)

type BankingRepository struct {
	cfg            *config.Config
	dbClientMaster postgre.DBInterface
	dbClientSlave  postgre.DBInterface
	queries        statementQueries
	httpClient     *http.Client
}

func NewBankingRepository(cfg *config.Config, masterClient, slaveClient postgre.DBInterface) ports.BankingRepositoryAdapter {
	repo := &BankingRepository{
		cfg:            cfg,
		dbClientMaster: masterClient,
		dbClientSlave:  slaveClient,
		queries:        statementQueries{},
	}
	err := repo.initDBSchema(context.TODO())
	if err != nil {
		log.Err(err)
	}

	err = repo.prepareStatements()
	if err != nil {
		log.Err(err)
	}
	return repo
}
