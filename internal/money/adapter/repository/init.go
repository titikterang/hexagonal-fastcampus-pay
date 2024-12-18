package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/postgre"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	mongo2 "go.mongodb.org/mongo-driver/v2/mongo"
)

type MoneyRepository struct {
	cfg                  *config.Config
	dbClientMaster       postgre.DBInterface
	dbClientSlave        postgre.DBInterface
	dbClient             mongo.DBInterface
	DB                   *mongo2.Database
	queries              statementQueries
	redisClient          *redis.Client
	kafkaClient          kafka.KafkaClientInterface
	topicValidationReply string
	topicMoneyDLQ        string
}

func NewMoneyRepository(cfg *config.Config,
	redisClient *redis.Client,
	dbClient mongo.DBInterface,
	masterClient,
	slaveClient postgre.DBInterface,
	kafkaClient kafka.KafkaClientInterface) ports.MoneyRepositoryAdapter {
	repo := &MoneyRepository{
		cfg:            cfg,
		dbClientMaster: masterClient,
		dbClientSlave:  slaveClient,
		queries:        statementQueries{},
		redisClient:    redisClient,
		kafkaClient:    kafkaClient,
		dbClient:       dbClient,
		DB:             dbClient.Database(cfg.MongoDB.DBName),
	}
	err := repo.initDBSchema(context.TODO())
	if err != nil {
		log.Err(err)
	}

	err = repo.prepareStatements()
	if err != nil {
		log.Err(err)
	}

	for k, v := range cfg.Kafka.ProducerTopics {
		switch k {
		case model.TopicTrxValidate:
			repo.topicValidationReply = v
		case model.TopicMoneyDLQ:
			repo.topicMoneyDLQ = v
		default:

		}
	}

	return repo
}

func (r *MoneyRepository) CloseClient() {
	r.kafkaClient.Close()
}
