package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
)

type Handler struct {
	config            *config.Config
	settlementService ports.SettlementServiceAdapter
}

type ConsumerHandler struct {
	config            *config.Config
	settlementService ports.SettlementConsumerAdapter
	client            kafka.KafkaClientInterface
	topicTransfer     string
	topicPayment      string
}

func NewHandler(cfg *config.Config, adapter ports.SettlementServiceAdapter) (*Handler, error) {
	return &Handler{
		config:            cfg,
		settlementService: adapter,
	}, nil
}

func NewConsumer(cfg *config.Config, client kafka.KafkaClientInterface, adapter ports.SettlementConsumerAdapter) (*ConsumerHandler, error) {
	handler := &ConsumerHandler{
		config:            cfg,
		settlementService: adapter,
		client:            client,
	}

	for k, v := range cfg.Kafka.ConsumerTopics {
		switch k {
		case model.TopicTransfer:
			handler.topicTransfer = v
		case model.TopicPayment:
			handler.topicPayment = v
		default:
			panic("unhandled default case")
		}
	}

	if handler.settlementService != nil {
		handler.settlementService.InitWorkerPool()
	}

	return handler, nil
}
