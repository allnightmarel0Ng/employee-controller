package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/storage/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/employee-controller/internal/logger"
	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
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

func (s *StorageHandler) HandleKafka() {
	logger.Debug("HandleKafka: start")
	defer logger.Debug("HandleKafka: end")

	for {
		msg, err := s.consumer.Consume(time.Second)
		if err == nil {
			readableMsg := string(msg.Value)
			logger.Trace("message on %s: %s", msg.TopicPartition, readableMsg)
			err = s.useCase.ProcessMessage(s.ctx, readableMsg)
			if err != nil {
				logger.Warning("unable to process message %s: %s", readableMsg, err.Error())
			}
		} else if !err.(confluentkafka.Error).IsTimeout() {
			logger.Warning("consumer error: %s", err.Error())
		}
	}
}

func (s *StorageHandler) GetEmployeeByIP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("GetEmployeeByIP: start")
	defer logger.Debug("GetEmployeeByIP: end")

	IP := chi.URLParam(r, "IP")
	employee, err := s.useCase.GetEmployee(s.ctx, IP)
	if err != nil {
		logger.Trace("employee not found")
		http.Error(w, "Employee not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
