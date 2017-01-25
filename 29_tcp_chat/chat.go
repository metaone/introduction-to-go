package main

import "net"


type Chat struct {
	clients map[string]*Client
	rooms map[string]*ChatRoom
}


func (chat *Chat) Join(conn net.Conn) {
	client := NewClient(conn)
	chat.clients[client.id] = client
	chat.rooms[DEFAULT_ROOM].clients[client.id] = client
}

func (chat *Chat) Leave(client *Client) {
	client.output <- Message{fromId: client.id, text: "/exit"}
	delete(chat.rooms[client.room].clients, client.id)
	delete(chat.clients, client.id)
	client.conn.Close()
}
