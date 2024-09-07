package usecase

import "github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"

type CollectorUseCase interface {
	ProcessMessage(topic string, msg []byte) error
}

type collectorUseCase struct {
	producer *kafka.Producer
}

func NewCollectorUseCase(producer *kafka.Producer) CollectorUseCase {
	return &collectorUseCase{
		producer: producer,
	}
}

func (c *collectorUseCase) ProcessMessage(topic string, msg []byte) error {
	return c.producer.Produce(topic, msg)
}
