package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type MessageType int

const (
	Info MessageType = iota
	Activity
	Screenshot
)

type Message struct {
	Type MessageType
	Data interface{}
}

type Employee struct {
	Domain       string
	Machine      string
	User         string
	IP           net.Addr
	LastActivity time.Time
}

func (e Employee) Print() {
	fmt.Printf("%s %s %s %s %s\n", e.Domain, e.Machine, e.User, e.IP.String(), e.LastActivity.String())
}

func server(channel <-chan Message) {
	employees := make(map[string]*Employee)

	var mu sync.Mutex
	go func() {
		for {
			message := <-channel
			mu.Lock()

			switch message.Type {
			case Info:
				employee := message.Data.(Employee)
				employees[employee.IP.String()] = &employee
			case Activity:
				employees[message.Data.(string)].LastActivity = time.Now()
			case Screenshot:

			}
			mu.Unlock()
		}
	}()
}

func handleInfo(msg string, channel chan<- Message, addr net.Addr) {
	tokens := strings.Split(msg, " ")
	if len(tokens) != 3 {
		return
	}

	channel <- Message{
		Type: Info,
		Data: Employee{
			Domain:       tokens[0],
			Machine:      tokens[1],
			User:         tokens[2],
			IP:           addr,
			LastActivity: time.Now(),
		},
	}
}

func handleActivity(channel chan<- Message, addr net.Addr) {
	channel <- Message{
		Type: Activity,
		Data: addr.String(),
	}
}

func handleConnection(conn net.Conn, channel chan<- Message) {
	remoteAddr := conn.RemoteAddr()
	log.Printf("%s connected", remoteAddr)

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
		msg := string(raw)
		log.Printf("%s: %s", remoteAddr, msg)

		if msg[:4] == "INFO" {
			handleInfo(msg[5:], channel, remoteAddr)
		} else if msg[:8] == "ACTIVITY" {
			handleActivity(channel, remoteAddr)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Wrong amount of arguments")
	}

	listener, err := net.Listen("tcp", "localhost:"+os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	channel := make(chan Message)
	go server(channel)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConnection(connection, channel)
	}
}
