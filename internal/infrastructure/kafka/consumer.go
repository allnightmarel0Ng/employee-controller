package kafka

import (
	"time"

	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	consumer *confluentkafka.Consumer
}

func NewConsumer(broker string, group string) (*Consumer, error) {
	c, err := confluentkafka.NewConsumer(&confluentkafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: c,
	}, nil
}

func (c *Consumer) SubscribeTopics(topics []string) error {
	return c.consumer.SubscribeTopics(topics, nil)
}

func (c *Consumer) Consume(timeout time.Duration) (*confluentkafka.Message, error) {
	return c.consumer.ReadMessage(timeout)
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
