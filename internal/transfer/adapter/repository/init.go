package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"net/http"
	"time"
)

type TransferRepository struct {
	cfg                *config.Config
	dbClientMaster     postgre.DBInterface
	dbClientSlave      postgre.DBInterface
	queries            statementQueries
	redisClient        *redis.Client
	httpClient         *http.Client
	kafkaClient        kafka.KafkaClientInterface
	topicProducerMoney string
	topicProducerBank  string
}

func NewTransferRepository(cfg *config.Config,
	redisClient *redis.Client,
	masterClient, slaveClient postgre.DBInterface,
	kafkaClient kafka.KafkaClientInterface) ports.TransferServiceRepositoryAdapter {

	repo := &TransferRepository{
		cfg:            cfg,
		dbClientMaster: masterClient,
		dbClientSlave:  slaveClient,
		queries:        statementQueries{},
		redisClient:    redisClient,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     10 * time.Second,
				MaxIdleConns:        100,
			},
		},
		kafkaClient:        kafkaClient,
		topicProducerMoney: "",
		topicProducerBank:  "",
	}

	for k, v := range cfg.Kafka.ConsumerTopics {
		switch k {
		case model.MoneyTopicKey:
			repo.topicProducerMoney = v
		case model.BankingTopicKey:
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
