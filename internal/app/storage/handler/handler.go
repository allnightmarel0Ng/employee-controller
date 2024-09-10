package handler

import (
	"context"
	"encoding/json"
	"errors"
	confluentkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/storage/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
)

type StorageHandler struct {
	useCase  usecase.StorageUseCase
	consumer *kafka.Consumer
	ctx      context.Context
}

func NewStorageHandler(useCase usecase.StorageUseCase, consumer *kafka.Consumer, ctx context.Context) StorageHandler {
	return StorageHandler{
		useCase:  useCase,
		consumer: consumer,
		ctx:      ctx,
	}
}

func (s *StorageHandler) HandleConsumer() {
	for {
		msg, err := s.consumer.Consume(time.Second)
		if err == nil {
			readableMsg := string(msg.Value)
			log.Printf("Message on %s: %s", msg.TopicPartition, readableMsg)
			err = s.useCase.ProcessMessage(s.ctx, readableMsg)
			if err != nil {
				log.Printf("Unable to process message %s: %s", readableMsg, err.Error())
			}
		} else {
			var kafkaErr confluentkafka.Error
			ok := errors.As(err, &kafkaErr)
			if ok && kafkaErr.Code() != confluentkafka.ErrTimedOut {
				log.Printf("Consumer error: %s", err.Error())
			}
		}
	}
}

func (s *StorageHandler) GetEmployeeByIP(w http.ResponseWriter, r *http.Request) {
	IP := chi.URLParam(r, "IP")
	employee, err := s.useCase.GetEmployee(s.ctx, IP)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
