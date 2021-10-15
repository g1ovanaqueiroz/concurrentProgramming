// a TCP server that periodically writes the time
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	count := 0

	go sendAuditLog()

	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		count = count + 1
		fmt.Printf("Number of Active %d\n", count)
		go handleConn(conn, count)
	}
}

func sendAuditLog(chClient1 <-chan string, chClient2 <-chan string) {
	// connect to yet another server to send audit log
	// every time a msg is sent back to a client

	//eat your own dog food
	timerCh := time.Tick(10 * time.Second)
	maxTries := 5
	tries := 0

	for {
		select {
		case msg := <-chClient1:
			sendToServer(msg)
		case msg := <-chClient2:
			sendToServer(msg)
		case <-timerCh:
			tries = tries + 1
			if tries >= maxTries:
				fmt.Println("timeout")
		}

	}

	//timestamp str, int counter
}

func handleConn(c net.Conn, count int) {
	for {
		_, err := io.WriteString(c, time.Now().String()+"\n")

		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func sendToServer(msg string) {
	//code sending msg to server
}
