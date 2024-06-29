package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	fmt.Println("Starting Echo Server...")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	defer listen.Close()
	fmt.Println("Server listening on", HOST+":"+PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %s", err)
			}
			break
		}

		_, err = conn.Write(buffer[:n])
		if err != nil {
			log.Printf("Error writing to connection: %s", err)
			break
		}
	}
}
