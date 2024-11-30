package handler

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"github.com/twmb/franz-go/pkg/kgo"
)

func (c *ConsumerHandler) StartConsumer() {
	for {
		ctx := context.Background()
		fetches := c.client.PollRecords(ctx, c.config.Kafka.MaxPollRecord)
		if fetches.IsClientClosed() {
			return
		}

		fetches.EachError(func(t string, p int32, err error) {
			log.Error().Msgf("fetch failed for topic %s partition %d", t, p)
		})

		fetches.EachRecord(func(r *kgo.Record) {
			// log request id from header
			// handle message based on topic
			switch r.Topic {
			case c.topicTransactionValidate:
				var data types.TransactionValidateInfo
				err := json.Unmarshal(r.Value, &data)
				if err != nil {
					log.Error().Msgf("Unmarshal HandleTransactionValidation failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}
				err = c.moneyService.HandleTransactionValidation(r.Context, data)
				if err != nil {
					log.Error().Msgf("HandleTransactionValidation failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}
			}
			c.client.MarkCommitRecords(r)
		})
		c.client.AllowRebalance()
	}
}
