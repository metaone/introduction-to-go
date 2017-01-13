//package main
//
//import (
//	"fmt"
//	"bufio"
//	"net"
//	//"io"
//)
//
//type User struct {
//	in chan string
//	out chan string
//	reader *bufio.Reader
//	writer *bufio.Writer
//}
//
//type ChatRoom struct {
//	users []*User
//	in chan string
//	out chan string
//}
//
//func (room *ChatRoom) Join(user *User) {
//
//}
//
//func main() {
//	fmt.Print("here\n")
//
//	server, _ := net.Listen("tcp", ":11111")
//	defer server.Close()
//
//	for {
//		conn, _ := server.Accept()
//		go handle(conn)
//	}
//}
//
//func handle(conn net.Conn) {
//	data := make([]byte, 1024)
//	read, _ := conn.Read(data)
//	fmt.Print(&read, "\n")
//}

package main

import (
	"bufio"
	"net"
	"fmt"
)

type Client struct {
	incoming chan string
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.incoming <- line
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

	client := &Client{
		incoming: make(chan string),
		outgoing: make(chan string),
		reader: reader,
		writer: writer,
	}

	client.Listen()

	return client
}

type ChatRoom struct {
	clients []*Client
	joins chan net.Conn
	incoming chan string
	outgoing chan string
}

func (chatRoom *ChatRoom) Broadcast(data string) {
	for _, client := range chatRoom.clients {
		client.outgoing <- data
	}
}

func (chatRoom *ChatRoom) Join(connection net.Conn) {
	client := NewClient(connection)
	chatRoom.clients = append(chatRoom.clients, client)
	go func() { for { chatRoom.incoming <- <-client.incoming } }()
}

func (chatRoom *ChatRoom) Listen() {
	go func() {
		for {
			select {
			case data := <-chatRoom.incoming:
				chatRoom.Broadcast(data)
			case conn := <-chatRoom.joins:
			fmt.Print("join room\n", conn.LocalAddr(), "\n", conn.RemoteAddr(), "\n")
				chatRoom.Join(conn)
			}
		}
	}()
}

func NewChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		clients: make([]*Client, 0),
		joins: make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	chatRoom.Listen()

	return chatRoom
}

func main() {
	chatRoom := NewChatRoom()

	listener, _ := net.Listen("tcp", ":6666")

	for {
		conn, _ := listener.Accept()
		chatRoom.joins <- conn
	}
}
