package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"net/http"
	"time"
)

const (
	moneyTopicKey   = "money_service"
	bankingTopicKey = "banking_service"
)

type TransferRepository struct {
	cfg                *config.Config
	dbClientMaster     postgre.DBInterface
	dbClientSlave      postgre.DBInterface
	queries            statementQueries
	httpClient         *http.Client
	kafkaClient        kafka.KafkaClientInterface
	topicProducerMoney string
	topicProducerBank  string
}

func NewTransferRepository(cfg *config.Config,
	masterClient, slaveClient postgre.DBInterface,
	kafkaClient kafka.KafkaClientInterface) ports.TransferServiceRepositoryAdapter {

	repo := &TransferRepository{
		cfg:            cfg,
		dbClientMaster: masterClient,
		dbClientSlave:  slaveClient,
		queries:        statementQueries{},
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     10 * time.Second,
				MaxIdleConns:        100,
			},
		},
		kafkaClient: kafkaClient,
	}

	for k, v := range cfg.Kafka.ConsumerTopics {
		switch k {
		case moneyTopicKey:
			repo.topicProducerMoney = v
		case bankingTopicKey:
			repo.topicProducerBank = v
		}
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
