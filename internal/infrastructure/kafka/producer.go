package kafka

import (
	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	producer *confluentkafka.Producer
}

func NewProducer(broker string) (*Producer, error) {
	p, err := confluentkafka.NewProducer(&confluentkafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: p,
	}, nil
}

func (p *Producer) Produce(topic string, message []byte) error {
	return p.producer.Produce(&confluentkafka.Message{
		TopicPartition: confluentkafka.TopicPartition{Topic: &topic, Partition: confluentkafka.PartitionAny},
		Value:          message,
	}, nil)
}

func (p *Producer) Close() {
	p.producer.Close()
}
