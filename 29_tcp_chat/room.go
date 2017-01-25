package main


type ChatRoom struct {
	clients map[string]*Client
	input chan Message
}

func (room *ChatRoom) Broadcast(msg Message) {
	for _, client := range room.clients {
		if client.id != msg.fromId {
			client.output <- msg
		}
	}
}

func (room *ChatRoom) Loop() {
	for {
		msg := <-room.input
		room.Broadcast(msg)
	}
}

func NewRoom() *ChatRoom {
	room := &ChatRoom{
		clients: make(map[string]*Client),
		input: make(chan Message),
	}

	go room.Loop()

	return room
}
