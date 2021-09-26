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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	for {
		_, err := io.WriteString(c, time.Now().String()+"\n")

		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
