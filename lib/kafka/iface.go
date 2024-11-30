package kafka

import (
	"context"
	"github.com/twmb/franz-go/pkg/kgo"
)

// KafkaProducerInterface ...
type KafkaProducerInterface interface {
	Produce(ctx context.Context, topic string, key []byte, value []byte) error
	ProduceSync(ctx context.Context, records ...*kgo.Record) error
}

// KafkaClientInterface interface for kafka handler
type KafkaClientInterface interface {
	ProduceSync(ctx context.Context, rs ...*kgo.Record) kgo.ProduceResults
	CommitRecords(ctx context.Context, rs ...*kgo.Record) error
	MarkCommitRecords(rs ...*kgo.Record)
	PollRecords(ctx context.Context, maxPollRecords int) kgo.Fetches
	PollFetches(ctx context.Context) kgo.Fetches
	AllowRebalance()
	Ping(ctx context.Context) error
	Close()
}
