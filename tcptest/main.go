package main

import (
	"fmt"
	"net"
)

import (
	"tcptest/models"
)

func handleConnection(conn net.Conn) {
	// Read data from the connection
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Print the received data
	fmt.Println("Received data:", string(buffer))

	// Send a response back to the client
	conn.Write([]byte("Hello from the server!"))

	// Close the connection when done
	conn.Close()
}

func main() {
	// Listen on a specific address and port
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on localhost:8080")

	// Accept and handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection concurrently (in a goroutine)
		go handleConnection(conn)
	}
}