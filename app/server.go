package main

import (
	"fmt"

	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		conn.Write([]byte("+PONG\r\n"))
	}
}
