package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
)

type Handler struct {
	config         *config.Config
	bankingService ports.BankingServiceAdapter
}

type ConsumerHandler struct {
	config            *config.Config
	bankingService    ports.BankingServiceConsumerAdapter
	client            kafka.KafkaClientInterface
	topicTransferBank string
	topicPaymentBank  string
}

func NewHandler(cfg *config.Config, adapter ports.BankingServiceAdapter) (*Handler, error) {
	return &Handler{
		config:         cfg,
		bankingService: adapter,
	}, nil
}

func NewConsumer(cfg *config.Config, client kafka.KafkaClientInterface, adapter ports.BankingServiceConsumerAdapter) (*ConsumerHandler, error) {
	handler := &ConsumerHandler{
		config:         cfg,
		bankingService: adapter,
		client:         client,
	}

	for k, v := range cfg.Kafka.ConsumerTopics {
		switch k {
		case model.TopicBankTransfer:
			handler.topicTransferBank = v
		case model.TopicBankPayment:
			handler.topicPaymentBank = v
		}
	}

	return handler, nil
}
