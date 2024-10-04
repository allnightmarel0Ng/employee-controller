package main

import (
	"context"
	"net/http"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/storage/handler"
	"github.com/allnightmarel0Ng/employee-controller/internal/app/storage/repository"
	"github.com/allnightmarel0Ng/employee-controller/internal/app/storage/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/config"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/redis"
	"github.com/allnightmarel0Ng/employee-controller/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	ctx := context.Background()

	redisClient := redis.NewClient("redis:"+conf.RedisPort, "", 0)
	defer redisClient.Close()

	repo := repository.NewStorageRepository(redisClient)

	consumer, err := kafka.NewConsumer("broker:"+conf.KafkaBroker, "activities")
	if err != nil {
		logger.Error("unable to start the kafka consumer: %s", err.Error())
	}
	defer consumer.Close()
	err = consumer.SubscribeTopics([]string{"events"})
	if err != nil {
		logger.Error("unable to subscribe consumer on topics: %s", err.Error())
	}

	useCase := usecase.NewStorageUseCase(repo)
	handle := handler.NewStorageHandler(useCase, consumer, ctx)

	go handle.HandleKafka()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/employee/{IP}", handle.GetEmployeeByIP)

	err = http.ListenAndServe(":"+conf.HTTPPort, router)
	if err != nil {
		logger.Error("unable to start http server: %s", err.Error())
	}
}
