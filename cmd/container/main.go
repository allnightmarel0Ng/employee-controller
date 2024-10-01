package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/container/handler"
	"github.com/allnightmarel0Ng/employee-controller/internal/app/container/repository"
	"github.com/allnightmarel0Ng/employee-controller/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/config"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/postgres"
	"github.com/allnightmarel0Ng/employee-controller/internal/logger"
	pb "github.com/allnightmarel0Ng/employee-controller/internal/protos/container"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.LogError("unable to load config: %s", err.Error())
	}

	db, err := postgres.NewDatabase(context.Background(),
		fmt.Sprintf("postgresql://%s:%s@postgres:%s/%s?sslmode=disable",
			conf.PostgresUser,
			conf.PostgresPassword,
			conf.PostgresPort,
			conf.PostgresName))
	if err != nil {
		logger.LogError("unable to connect to database: %s", err.Error())
	}
	defer db.Close()

	consumer, err := kafka.NewConsumer("broker:"+conf.KafkaBroker, "mygroup")
	if err != nil {
		logger.LogError("unable to connect to database: %s", err.Error())
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{"events"})
	if err != nil {
		logger.LogError("unable to subscribe kafka consumer on topic: %s", err.Error())
	}

	repo := repository.NewContainerRepository(db)
	useCase := usecase.NewContainerUseCase(repo)

	kafkaHandler := handler.NewContainerKafkaHandler(consumer, useCase)
	go kafkaHandler.HandleKafka()

	listener, err := net.Listen("tcp", ":"+conf.ContainerPort)
	if err != nil {
		logger.LogError("unable to listen on port %s: %s", conf.ContainerPort, err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterContainerServer(server, &handler.ContainerGRPCHandler{UseCase: useCase})

	err = server.Serve(listener)
	if err != nil {
		logger.LogError("unable to server on port %s: %s", conf.ContainerPort, err.Error())
	}
}
