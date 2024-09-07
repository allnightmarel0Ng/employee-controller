package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(broker string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.service": broker})
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: p,
	}, nil
}

func (p *Producer) Produce(topic string, message []byte) error {
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
}

func (p *Producer) Close() {
	p.producer.Close()
}
