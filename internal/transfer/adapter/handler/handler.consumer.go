package handler

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
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
			log.Info().Msgf("option %s vs %s", c.topicMoneyReply, c.topicBankReply)
			switch r.Topic {
			case c.topicMoneyReply:
				var data types.TransactionValidateReplyInfo
				err := json.Unmarshal(r.Value, &data)
				if err != nil {
					log.Error().Msgf("Unmarshal TransactionValidateReplyInfo failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}
				err = c.transferService.HandleTransactionValidationReply(r.Context, data)
				if err != nil {
					log.Error().Msgf("HandleTransactionValidationReply failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}
			case c.topicBankReply:
				var data types.TransferExternalBankReply
				err := json.Unmarshal(r.Value, &data)
				if err != nil {
					log.Error().Msgf("TransferExternalBankReply Unmarshal failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}
				err = c.transferService.HandleBankCallbackReply(r.Context, data)
				if err != nil {
					log.Error().Msgf("HandleBankCallbackReply failed for topic %s partition %d", r.Topic, r.Partition)
					return
				}
			}
			c.client.MarkCommitRecords(r)
		})
		c.client.AllowRebalance()
	}
}
