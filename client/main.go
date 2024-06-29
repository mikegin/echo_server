package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	HOST            = "127.0.0.1"
	PORT            = "8080"
	NUM_CONNECTIONS = 5
)

func handleRequest(conn net.Conn, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error sending data to server: %v\n", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error receiving data from server: %v\n", err)
		return
	}

	fmt.Printf("Received from server: %s\n", string(buffer[:n]))
}

func createClients() {
	var wg sync.WaitGroup

	for i := 0; i < NUM_CONNECTIONS; i++ {
		conn, err := net.Dial("tcp", HOST+":"+PORT)
		if err != nil {
			fmt.Printf("Connection %d failed: %v\n", i+1, err)
			return
		}

		message := fmt.Sprintf("Hello from client %d", i+1)

		wg.Add(1)
		go handleRequest(conn, message, &wg)
	}

	wg.Wait()
}

func main() {
	createClients()
}
