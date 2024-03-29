package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const pong = "+PONG\r\n"

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	for {
		input := make([]byte, 2048)
		_, err := c.Read(input)
		if err != nil {
			log.Fatal("could not read Conn")
		}

		c.Write([]byte(pong))
	}
}
