package main

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("INFO " + hostname + " " + currentUser.Username))
	if err != nil {
		os.Exit(1)
	}

	lastActivity := time.Now()

	for {
		if time.Since(lastActivity) > 30*time.Second {
			if robotgo.MouseSleep == 0 || robotgo.KeySleep == 0 {
				lastActivity = time.Now()
				conn.Write([]byte("ACTIVITY"))
			}
		}
		time.Sleep(time.Second)
	}
}
