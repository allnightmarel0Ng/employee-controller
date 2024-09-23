package handler

import (
	"fmt"
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
			log.Printf("Couldn't close the connection with %s: %s", remoteAddr, err.Error())
		}
	}(conn)

	run := true
	for run {
		raw := make([]byte, 1024)
		bytesRead, err := conn.Read(raw)

		var msg string
		if bytesRead == 0 {
			msg = fmt.Sprintf("DISCONNECTED %s", remoteAddr.String())
			log.Printf("%s disconnected", remoteAddr)
			run = false
		} else if err == nil {
			msg = fmt.Sprintf("%s %s", string(raw), remoteAddr.String())
		} else {
			log.Printf("Couldn't read message from %s: %s", remoteAddr, err.Error())
			continue
		}

		err = c.useCase.ProcessMessage("events", []byte(msg))
		if err != nil {
			log.Printf("Couldn't process message from %s: %s", remoteAddr, err.Error())
		}
	}
}
