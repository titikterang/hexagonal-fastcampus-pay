package kafka

import (
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type FastCampusPayConsumer struct {
	cfg    *config.Config
	client KafkaClientInterface
}

func InitKafkaClient(cfg *config.Config) (*FastCampusPayConsumer, error) {
	var (
		consumer  FastCampusPayConsumer
		err       error
		topicList []string
	)

	for _, v := range cfg.Kafka.ConsumerTopics {
		topicList = append(topicList, v)
	}

	consumer.cfg = cfg
	consumerOptions := []kgo.Opt{
		kgo.SeedBrokers(cfg.Kafka.Hosts...),
		kgo.ConsumerGroup(cfg.Kafka.ConsumerGroup),
		kgo.ConsumeTopics(topicList...),
		kgo.FetchIsolationLevel(kgo.ReadUncommitted()),
		kgo.AutoCommitMarks(),
		kgo.BlockRebalanceOnPoll(),
		kgo.AllowAutoTopicCreation(),
	}

	consumer.client, err = kgo.NewClient(consumerOptions...)
	if err != nil {
		log.Error().Msgf("kgo.NewClient err %#v ", err)
		return &consumer, err
	}

	return &consumer, nil
}

func (c *FastCampusPayConsumer) CloseClient() {
	c.client.Close()
}
