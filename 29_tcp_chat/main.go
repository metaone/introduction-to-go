package main

import (
	"net"
	"log"
)

var chat *Chat

const DEFAULT_ROOM = "global"

func main() {
	chat = &Chat{
		rooms: make(map[string]*ChatRoom),
		clients: make(map[string]*Client),
	}

	chat.rooms[DEFAULT_ROOM] = NewRoom()

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
