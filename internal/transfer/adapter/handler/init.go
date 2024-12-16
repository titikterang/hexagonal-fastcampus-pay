package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/transfer/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
)

type Handler struct {
	config          *config.Config
	transferService ports.TransferServiceAPIAdapter
}

type ConsumerHandler struct {
	config          *config.Config
	transferService ports.TransferServiceConsumerAdapter
	client          kafka.KafkaClientInterface
	topicMoneyReply string
	topicBankReply  string
}

func NewHandler(cfg *config.Config, adapter ports.TransferServiceAPIAdapter) (*Handler, error) {
	return &Handler{
		config:          cfg,
		transferService: adapter,
	}, nil
}

func NewConsumer(cfg *config.Config, client kafka.KafkaClientInterface, adapter ports.TransferServiceConsumerAdapter) (*ConsumerHandler, error) {
	handler := &ConsumerHandler{
		config:          cfg,
		transferService: adapter,
		client:          client,
	}

	for k, v := range cfg.Kafka.ConsumerTopics {
		switch k {
		case model.MoneyTopicKey:
			handler.topicMoneyReply = v
		case model.BankingTopicKey:
			handler.topicBankReply = v
		default:
			panic("unhandled default case")
		}
	}

	return handler, nil
}
