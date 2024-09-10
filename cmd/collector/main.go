package main

import (
	"log"
	"net"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/collector/handler"
	"github.com/allnightmarel0Ng/employee-controller/internal/app/collector/usecase"
	"github.com/allnightmarel0Ng/employee-controller/internal/config"
	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/kafka"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	producer, err := kafka.NewProducer("broker:" + conf.KafkaBroker)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer producer.Close()

	TCPHandler := handler.NewCollectorHandler(usecase.NewCollectorUseCase(producer))

	listener, err := net.Listen("tcp", ":"+conf.TCPPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Printf("Unable to close the listener: %s", err.Error())
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
			continue
		}
		go TCPHandler.HandleConnection(conn)
	}
}
