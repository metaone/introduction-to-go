package main

import (
	"bufio"
	"net"
	"strings"
)

const DEFAULT_ROOM = "global"

var chat *Chat

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
	id string
	username string
	room string
	input chan string
	output chan string
	reader *bufio.Reader
	writer *bufio.Writer
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')

		// check for a chat command
		if strings.HasPrefix(line, "/") {
			command := strings.Split(strings.TrimSpace(line), " ")
			switch command[0] {
			case "/setname":
				client.username = command[1]
			case "/showrooms":
				for key := range chat.rooms {
					client.Message(key + "\n")
				}
			case "/addroom":
				if _, ok := chat.rooms[command[1]]; ok {
					client.Message("Error: room already exists\n")
				} else {
					chat.rooms[command[1]] = CreateRoom()
				}
			case "/joinroom":
				if _, ok := chat.rooms[command[1]]; ok {
					delete(chat.rooms[client.room].clients, client.id)
					client.room = command[1]
					chat.rooms[client.room].clients[client.id] = client
				} else {
					client.Message("Error: room not exists\n")
				}
			default:
				client.Message(help)
			}
		// add client info and forward further
		} else {
			chat.rooms[client.room].input <- client.id + "\n" + client.username + ": " + line
		}
	}
}

func (client *Client) Write() {
	for data := range client.output {
		client.Message(data)
	}
}

func (client *Client) Message(msg string) {
	client.writer.WriteString(msg)
	client.writer.Flush()
}

func CreateClient(conn net.Conn) *Client {
	client := &Client{
		id: conn.RemoteAddr().String(),
		username: conn.RemoteAddr().String(),
		room: DEFAULT_ROOM,
		input: make(chan string),
		output: make(chan string),
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}

	go client.Read()
	go client.Write()

	return client
}


type ChatRoom struct {
	clients map[string]*Client
	input chan string
	output chan string
}

func (room *ChatRoom) Broadcast(data string) {
	parsed := strings.Split(data, "\n")

	for _, client := range room.clients {
		if client.id != parsed[0] {
			client.output <- parsed[1] + "\n"
		}
	}
}

func CreateRoom() *ChatRoom {
	room := &ChatRoom{
		clients: make(map[string]*Client),
		input: make(chan string),
		output: make(chan string),
	}

	go func() {
		for {
			msg := <-room.input
			room.Broadcast(msg)
		}
	}()

	return room
}

type Chat struct {
	join chan net.Conn
	clients map[string]*Client
	rooms map[string]*ChatRoom
}

func (chat *Chat) Start() {
	go func() {
		for {
			conn := <-chat.join
			client := CreateClient(conn)
			chat.clients[client.id] = client
			chat.rooms[DEFAULT_ROOM].clients[client.id] = client
		}
	}()
}

func main() {
	chat = &Chat{
		join: make(chan net.Conn),
		rooms: make(map[string]*ChatRoom),
		clients: make(map[string]*Client),
	}

	chat.rooms[DEFAULT_ROOM] = CreateRoom()
	chat.Start()

	// start TCP
	listener, _ := net.Listen("tcp", ":6666")
	for {
		conn, _ := listener.Accept()
		chat.join <- conn
	}
}
