package main

import (
	"bufio"
	"net"
	"strings"
	"fmt"
)

var help string = `
HELP
/comand params ...
list of commands:
	/setname username
	/showrooms
	/addroom name
	/joinroom name
`

type Client struct {
	address string
	username string
	room string
	incoming chan string
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')

		// check for a chat command
		if strings.HasPrefix(line, "/") {fmt.Print(1, "\n")
			command := strings.Split(strings.TrimSpace(line), " ")
			switch command[0] {
			case "/setname":
				client.username = command[1]
			case "/showrooms":
				for key := range ChatRooms {
					client.writer.WriteString(key + "\n")
				}
				client.writer.Flush()
			case "/addroom":
				if _, ok := ChatRooms[command[1]]; ok {
					client.writer.WriteString("Error: room already exists\n")
					client.writer.Flush()
				} else {
					ChatRooms[command[1]] = NewChatRoom()
				}
			case "/joinroom":
				if room, ok := ChatRooms[command[1]]; ok {
					delete(ChatRooms[client.room].clients, client.address)
					client.room = command[1]
					room.Join(client)
					//client.incoming <- client.username + ": " + "leave room"
				} else {
					client.writer.WriteString("Error: room not exists\n")
					client.writer.Flush()
				}
			default:
				client.writer.WriteString(help)
				client.writer.Flush()
			}
		// add client info and forward further
		} else {fmt.Print(2, "\n")
			client.incoming <- client.address + "\n" + client.username + ": " + line
		}
	}
}

func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)
	address := connection.RemoteAddr().String()

	client := &Client{
		address: address,
		username: address,
		room: "global",
		incoming: make(chan string),
		outgoing: make(chan string),
		reader: reader,
		writer: writer,
	}

	client.Listen()

	return client
}

type ChatRoom struct {
	clients map[string]*Client
	joins chan net.Conn
	incoming chan string
	outgoing chan string
}

var ChatRooms map[string]*ChatRoom = make(map[string]*ChatRoom)

func (chatRoom *ChatRoom) Broadcast(data string) {
	parsed := strings.Split(data, "\n")

	for _, client := range chatRoom.clients {
		if client.address != parsed[0] {
			client.outgoing <- parsed[1] + "\n"
		}
	}
}

func (chatRoom *ChatRoom) Join(client *Client) {
	chatRoom.clients[client.address] = client

	go func() {
		for {
			if _, ok := chatRoom.clients[client.address]; ok {
				chatRoom.incoming <- <-client.incoming
			} else {
				return
			}
		}
	}()
}

func (chatRoom *ChatRoom) Listen() {
	go func() {
		for {
			select {
			case data := <-chatRoom.incoming:
				chatRoom.Broadcast(data)
			case conn := <-chatRoom.joins:
				chatRoom.Join(NewClient(conn))
			}
		}
	}()
}

func NewChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		clients: make(map[string]*Client),
		joins: make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	chatRoom.Listen()

	return chatRoom
}

func main() {
	ChatRooms["global"] = NewChatRoom()

	listener, _ := net.Listen("tcp", ":6666")

	for {
		conn, _ := listener.Accept()
		ChatRooms["global"].joins <- conn
	}
}
