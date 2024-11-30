package repository

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

func (r *MoneyRepository) PublishMoneyValidationMessageReply(ctx context.Context, info types.TransactionValidateReplyInfo) error {
	py, err := json.Marshal(&info)
	if err != nil {
		log.Error().Msgf("err Marshal PublishMoneyValidationMessageReply %#v", err)
	}
	record := &kgo.Record{
		Key:       []byte(info.SourceAccountNumber),
		Value:     py,
		Timestamp: time.Now(),
		Topic:     r.topicValidationReply,
	}
	r.kafkaClient.ProduceSync(ctx, record)
	return err
}
