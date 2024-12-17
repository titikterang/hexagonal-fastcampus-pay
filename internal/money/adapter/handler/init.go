package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
)

type Handler struct {
	money.UnimplementedMoneyServiceServer
	config       *config.Config
	moneyService ports.MoneyServiceAdapter
}

type ConsumerHandler struct {
	config                   *config.Config
	moneyService             ports.MoneyServiceConsumerAdapter
	client                   kafka.KafkaClientInterface
	topicTransactionValidate string
	topicDLQ                 string
}

func NewHandler(cfg *config.Config, adapter ports.MoneyServiceAdapter) (*Handler, error) {
	return &Handler{
		config:       cfg,
		moneyService: adapter,
	}, nil
}

func NewConsumer(cfg *config.Config, client kafka.KafkaClientInterface, adapter ports.MoneyServiceConsumerAdapter) (*ConsumerHandler, error) {
	handler := &ConsumerHandler{
		config:       cfg,
		moneyService: adapter,
		client:       client,
	}

	for k, v := range cfg.Kafka.ConsumerTopics {
		switch k {
		case model.TopicTrxValidate:
			handler.topicTransactionValidate = v
		case model.TopicMoneyDLQ:
			handler.topicDLQ = v
		}
	}

	return handler, nil
}
