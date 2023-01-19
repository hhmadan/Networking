package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send data
	conn.Write([]byte("Hello, server!"))

	// Read response
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
