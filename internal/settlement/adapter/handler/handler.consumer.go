package handler

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"github.com/twmb/franz-go/pkg/kgo"
)

func (c *ConsumerHandler) CloseClient() {
	c.client.Close()
}

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
			log.Info().Msgf("new message from topic %s", r.Topic)
			switch r.Topic {
			case c.topicTransfer:
				var data types.TransferExternalBankReply
				err := json.Unmarshal(r.Value, &data)
				if err != nil {
					log.Error().Msgf("Unmarshal TransferExternalBankReply failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}

				c.settlementService.HandleSettlementQueue(model.SettlementPayload{
					TransactionID:  data.TransactionID,
					SettlementType: model.TransferSettlement,
				})
			case c.topicPayment:
				var data types.PaymentBankReply
				err := json.Unmarshal(r.Value, &data)
				if err != nil {
					log.Error().Msgf("PaymentBankReply Unmarshal failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}

				c.settlementService.HandleSettlementQueue(model.SettlementPayload{
					TransactionID:  data.TransactionID,
					SettlementType: model.PaymentSettlement,
				})
			}
			c.client.MarkCommitRecords(r)
		})
		c.client.AllowRebalance()
	}
}
