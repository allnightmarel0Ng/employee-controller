package handler

import (
	"time"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/employee-controller/internal/logger"
	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ContainerKafkaHandler struct {
	consumer *kafka.Consumer
	useCase  usecase.ContainerUseCase
}

func NewContainerKafkaHandler(consumer *kafka.Consumer, useCase usecase.ContainerUseCase) *ContainerKafkaHandler {
	return &ContainerKafkaHandler{
		consumer: consumer,
		useCase:  useCase,
	}
}

func (c *ContainerKafkaHandler) HandleKafka() {
	logger.Debug("HandleKafka: start")
	defer logger.Debug("HandleKafka: end")

	for {
		msg, err := c.consumer.Consume(time.Second)
		if err == nil {
			msg := string(msg.Value)
			logger.Trace("got new message: %s", msg)
			err = c.useCase.ProcessKafkaMessage(msg)
			if err != nil {
				logger.Warning("got an error while storing the new message: %s", err.Error())
			} else {
				logger.Trace("new message was processed successfully successfully")
			}
		} else if !err.(confluentkafka.Error).IsTimeout() {
			logger.Warning("consumer error: %s", err.Error())
		}
	}
}
