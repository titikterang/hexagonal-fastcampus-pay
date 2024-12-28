package repository

import (
	"github.com/eapache/go-resiliency/breaker"
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/membership"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
	mongo2 "go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"time"
)

type circuitBreaker struct {
	breakerAccount *breaker.Breaker
	breakerBalance *breaker.Breaker
}

type PaymentRepository struct {
	cfg                *config.Config
	dbClient           mongo.DBInterface
	cb                 circuitBreaker
	DB                 *mongo2.Database
	redisClient        *redis.Client
	httpClient         *http.Client
	kafkaClient        kafka.KafkaClientInterface
	topicProducerMoney string
	topicProducerBank  string
	// grpc client
	membershipClient membership.MembershipServiceClient
	moneyClient      money.MoneyServiceClient
}

func NewPaymentRepository(cfg *config.Config,
	redisClient *redis.Client,
	dbClient mongo.DBInterface,
	kafkaClient kafka.KafkaClientInterface,
	membershipClient membership.MembershipServiceClient,
	moneyClient money.MoneyServiceClient) ports.PaymentRepositoryAdapter {

	repo := &PaymentRepository{
		cfg:         cfg,
		redisClient: redisClient,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     10 * time.Second,
				MaxIdleConns:        100,
			},
		},
		dbClient:           dbClient,
		DB:                 dbClient.Database(cfg.MongoDB.DBName),
		kafkaClient:        kafkaClient,
		topicProducerMoney: "",
		topicProducerBank:  "",
		// grpc client
		membershipClient: membershipClient,
		moneyClient:      moneyClient,
		cb: circuitBreaker{
			breakerAccount: breaker.New(5, 5, 3*time.Second),
			breakerBalance: breaker.New(5, 5, 3*time.Second),
		},
	}

	for k, v := range cfg.Kafka.ProducerTopics {
		switch k {
		case model.MoneyTopicKey:
			repo.topicProducerMoney = v
		case model.BankingTopicKey:
			repo.topicProducerBank = v
		default:

		}
	}

	return repo
}

func (r *PaymentRepository) CloseClient() {
	r.kafkaClient.Close()
}
