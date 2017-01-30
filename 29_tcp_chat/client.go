package main

import (
	"net"
	"bufio"
	"io"
	"log"
	"strings"
)

type Client struct {
	id string
	username string
	room string
	chat Chat
	conn net.Conn
	input chan string
	output chan Message
	reader *bufio.Reader
	writer *bufio.Writer
}

var help string = `
HELP

/comand params ...

list of commands:
/setname username
/showrooms
/addroom name
/joinroom name
/exit
`

func (client *Client) Read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err == io.EOF {
			client.chat.Leave(client)
			return
		} else if err != nil {
			log.Print(err)
			continue
		}

		if strings.HasPrefix(line, "/") {
			command := strings.Split(strings.TrimSpace(line), " ")
			switch command[0] {
			case "/setname":
				client.username = command[1]
			case "/showrooms":
				for key := range client.chat.rooms {
					client.Print(Message{text: key + "\n"})
				}
			case "/addroom":
				if _, ok := client.chat.rooms[command[1]]; ok {
					client.Print(Message{text: "Error: room already exists\n"})
				} else {
					client.chat.rooms[command[1]] = NewRoom()
				}
			case "/joinroom":
				if _, ok := client.chat.rooms[command[1]]; ok {
					if client.room != command[1] {
						delete(client.chat.rooms[client.room].clients, client.id)
						client.room = command[1]
						client.chat.rooms[client.room].clients[client.id] = client
					}
				} else {
					client.Print(Message{text: "Error: room not exists\n"})
				}
			case "/exit":
				client.chat.Leave(client)
				return
			default:
				client.Print(Message{text: help})
			}
		} else {
			client.chat.rooms[client.room].input <- Message{
				fromId: client.id,
				from: client.username,
				text: line,
			}
		}
	}
}

func (client *Client) Write() {
	for msg := range client.output {
		if msg.fromId == client.id && msg.text == "/exit" {
			return
		}
		client.Print(msg)
	}
}

func (client *Client) Print(msg Message) {
	text := msg.from
	if text != "" {
		text += ": "
	}
	text += msg.text

	client.writer.WriteString(text)
	client.writer.Flush()
}


func NewClient(conn net.Conn, chat Chat) *Client {
	address := conn.RemoteAddr().String()

	client := &Client{
		id: address,
		username: address,
		room: "global",
		chat: chat,
		conn: conn,
		input: make(chan string),
		output: make(chan Message),
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}

	go client.Read()
	go client.Write()

	client.Print(Message{text: help})

	return client
}
