package repository

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

func (r *PaymentRepository) PublishPaymentValidateRequest(ctx context.Context, info types.TransactionValidateInfo) error {
	log.Info().Msgf("PublishPaymentValidateRequest -> " + r.topicProducerMoney)
	py, err := json.Marshal(&info)
	if err != nil {
		log.Error().Msgf("err Marshal PublishPaymentValidateRequest %#v", err)
	}
	record := &kgo.Record{
		Key:       []byte(info.SourceAccountNumber),
		Value:     py,
		Timestamp: time.Now(),
		Topic:     r.topicProducerMoney,
	}
	r.kafkaClient.ProduceSync(ctx, record)
	return nil
}

func (r *PaymentRepository) PublishPaymentBankingRequest(ctx context.Context, payload types.PaymentBankExecution) error {
	return nil
}
