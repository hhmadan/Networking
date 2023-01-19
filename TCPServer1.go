package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :8080...")

	for {
		// Wait for a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Echo incoming data
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		data := buffer[:n]
		fmt.Println("Received:", string(data))
		time.Sleep(10 * time.Second) // 5 seconds delay

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
