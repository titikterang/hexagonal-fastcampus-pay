package repository

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

func (r *BankingRepository) PublishPaymentCallbackReply(ctx context.Context, data types.PaymentBankReply) error {
	log.Info().Msgf("PublishPaymentCallbackReply -> " + r.topicProducerPayment)
	val, err := json.Marshal(&data)
	if err != nil {
		log.Error().Msgf("err Marshal PublishPaymentCallbackReply %#v", err)
	}
	record := &kgo.Record{
		Value:     val,
		Timestamp: time.Now(),
		Topic:     r.topicProducerPayment,
	}
	r.kafka.ProduceSync(ctx, record)
	return nil
}
