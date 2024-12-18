package handler

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/kafka"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
)

type Handler struct {
	payment.UnimplementedPaymentServiceServer
	config         *config.Config
	paymentService ports.PaymentServiceAdapter
}

type ConsumerHandler struct {
	config          *config.Config
	paymentService  ports.PaymentConsumerAdapter
	client          kafka.KafkaClientInterface
	topicMoneyReply string
	topicBankReply  string
}

func NewHandler(cfg *config.Config, adapter ports.PaymentServiceAdapter) (*Handler, error) {
	return &Handler{
		config:         cfg,
		paymentService: adapter,
	}, nil
}

func NewConsumer(cfg *config.Config, client kafka.KafkaClientInterface, adapter ports.PaymentConsumerAdapter) (*ConsumerHandler, error) {
	handler := &ConsumerHandler{
		config:         cfg,
		paymentService: adapter,
		client:         client,
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
