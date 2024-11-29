package repository

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

func (r *TransferRepository) PublishTransferValidateRequest(ctx context.Context, info types.TransactionValidateInfo) error {
	py, err := json.Marshal(&info)
	if err != nil {
		log.Error().Msgf("err Marshal PublishTransferValidateRequest %#v", err)
	}
	record := &kgo.Record{
		Key:       []byte(info.SourceAccountNumber),
		Value:     py,
		Timestamp: time.Now(),
		Topic:     r.topicProducerMoney,
	}
	r.kafkaClient.ProduceSync(ctx, record)
	return err
}

func (r *TransferRepository) PublishTransferBankingRequest(ctx context.Context, payload types.TransferExternalBankPayload) error {
	py, err := json.Marshal(&payload)
	if err != nil {
		log.Error().Msgf("err Marshal PublishTransferValidateRequest %#v", err)
	}
	record := &kgo.Record{
		Key:       []byte(payload.TransactionInfo.SourceAccountNumber),
		Value:     py,
		Timestamp: time.Now(),
		Topic:     r.topicProducerBank,
	}
	r.kafkaClient.ProduceSync(ctx, record)
	return err
}
