package main

import (
	"net"
	"log"
)

func main() {
	chat := NewChat()

	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go chat.Join(conn)
	}
}
