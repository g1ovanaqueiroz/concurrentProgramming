// a TCP server that periodically writes the time
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
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
