package repository

import (
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"net/http"
)

type TransferRepository struct {
	cfg                *config.Config
	redisClient        *redis.Client
	httpClient         *http.Client
	kafkaClient        kafka.KafkaClientInterface
	topicProducerMoney string
	topicProducerBank  string
}
