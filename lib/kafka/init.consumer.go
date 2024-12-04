package kafka

import (
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type ClientType int

const (
	TypeConsumerClient ClientType = iota
	TypeProducerClient
)

func InitKafkaClient(cfg *config.Config, clientType ClientType) (KafkaClientInterface, error) {
	var (
		err       error
		topicList []string
	)

	for _, v := range cfg.Kafka.ConsumerTopics {
		topicList = append(topicList, v)
	}

	consumerOptions := []kgo.Opt{
		kgo.SeedBrokers(cfg.Kafka.Hosts...),
	}

	if clientType == TypeConsumerClient {
		consumerOptions = append(consumerOptions, []kgo.Opt{
			kgo.ConsumerGroup(cfg.Kafka.ConsumerGroup),
			kgo.ConsumeTopics(topicList...),
			kgo.FetchIsolationLevel(kgo.ReadUncommitted()),
			kgo.AutoCommitMarks(),
			kgo.BlockRebalanceOnPoll(),
			kgo.AllowAutoTopicCreation(),
		}...)
	}

	client, err := kgo.NewClient(consumerOptions...)
	if err != nil {
		log.Error().Msgf("kgo.NewClient err %#v ", err)
		return client, err
	}

	return client, nil
}
