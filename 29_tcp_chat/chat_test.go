package main

import (
	"testing"
	"net"
	"log"
)

func TestChat_Join(t *testing.T) {
	chat := &Chat{
		rooms: make(map[string]*ChatRoom),
		clients: make(map[string]*Client),
	}

	_ = chat

	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Print(err)
	}

	chat.Join(conn)

	conn.Close()
}

func TestChat_Leave(t *testing.T) {

}
