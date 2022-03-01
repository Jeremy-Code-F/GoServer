package gofarm

import (
	"fmt"
	"log"
	"net"
)

func Serve() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port 8080")
	for {
		fmt.Println("Waiting for connection")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Println("handling connection")
	msg := "HTTP/1.1 200 OK \n\n Hello, World!"
	_, err := connection.Write([]byte(msg))

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
}
