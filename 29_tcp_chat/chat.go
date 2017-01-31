package main

import (
	"net"
)


type Chat struct {
	clients map[string]*Client
	rooms map[string]*ChatRoom
}


func (chat *Chat) Join(conn net.Conn) {
	client := NewClient(conn, *chat)
	chat.clients[client.id] = client
	chat.rooms["global"].clients[client.id] = client
}

func (chat *Chat) Leave(client *Client) {
	close(client.output)
	client.conn.Close()
	delete(chat.rooms[client.room].clients, client.id)
	delete(chat.clients, client.id)
}

func NewChat() *Chat {
	chat := &Chat{
		rooms: make(map[string]*ChatRoom),
		clients: make(map[string]*Client),
	}

	chat.rooms["global"] = NewRoom()

	return chat
}
