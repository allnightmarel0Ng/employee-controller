package handler

import (
	"log"
	"net"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/collector/usecase"
)

type CollectorHandler struct {
	useCase usecase.CollectorUseCase
}

func NewCollectorHandler(useCase usecase.CollectorUseCase) *CollectorHandler {
	return &CollectorHandler{
		useCase: useCase,
	}
}

func (c *CollectorHandler) HandleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	log.Printf("%s connected", remoteAddr)

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Couldn't close the connection with %s: %s", remoteAddr, err)
		}
	}(conn)

	for {
		raw := make([]byte, 1024)
		bytesRead, err := conn.Read(raw)
		if bytesRead == 0 {
			log.Printf("%s disconnected", remoteAddr)
			break
		}
		if err != nil {
			log.Printf("Couldn't read message from %s: %s", remoteAddr, err)
			continue
		}

		err = c.useCase.ProcessMessage("event", raw)
		if err == nil {
			log.Printf("Couldn't process message from %s: %s", remoteAddr, err)
		}
	}
}
