package handler

import (
	"log"
	"time"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ContainerKafkaHandler struct {
	consumer *kafka.Consumer
	useCase  usecase.ContainerUseCase
}

func (c *ContainerKafkaHandler) HandleKafka() {
	for {
		msg, err := c.consumer.Consume(time.Second)
		if err == nil {
			msg := string(msg.Value)
			log.Printf("got new message: %s", msg)
			err = c.useCase.ProcessKafkaMessage(msg)
			if err != nil {
				log.Printf("got an error while storing the new message: %s", err.Error())
			} else {
				log.Printf("new message was processed successfully successfully")
			}
		} else if !err.(confluentkafka.Error).IsTimeout() {
			log.Printf("consumer error: %s", err.Error())
		}
	}
}
