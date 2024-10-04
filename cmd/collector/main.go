package main

import (
	"net"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/collector/handler"
	"github.com/allnightmarel0Ng/employee-controller/internal/app/collector/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/config"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/employee-controller/internal/logger"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	producer, err := kafka.NewProducer("broker:" + conf.KafkaBroker)
	if err != nil {
		logger.Error("unable to establish the producer: %s", err.Error())
	}
	defer producer.Close()

	TCPHandler := handler.NewCollectorHandler(usecase.NewCollectorUseCase(producer))

	listener, err := net.Listen("tcp", ":"+conf.TCPPort)
	if err != nil {
		logger.Error("unable to create listener: %s", err.Error())
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			logger.Error("unable to close the listener: %s", err.Error())
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("error accepting connection: %s", err.Error())
			continue
		}
		go TCPHandler.HandleConnection(conn)
	}
}
