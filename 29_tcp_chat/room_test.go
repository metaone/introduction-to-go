package main

import (
	"testing"
	//"net"
)

func TestChatRoom_Broadcast(t *testing.T) {
	//chat := NewChat()
	//
	//length := len(chat.rooms["global"].clients)
	//if length != 0 {
	//	t.Fatalf("got client count equal to '%q', expected '0'", length)
	//}
	//
	//listener, err := net.Listen("tcp", ":6666")
	//if err != nil {
	//	t.Fatal(err)
	//}
	////defer listener.Close()
	//
	//dial1, err := net.Dial("tcp", ":6666")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//_ = dial1
	////defer dial1.Close()
	//
	////
	////dial2, err := net.Dial("tcp", ":6666")
	////if err != nil {
	////	t.Fatal(err)
	////}
	////defer dial2.Close()
	////
	//conn1, err := listener.Accept()
	//if err != nil {
	//	t.Fatal(err)
	//}
	////defer conn1.Close()
	//
	//chat.Join(conn1)
	//length = len(chat.clients)
	//if length != 1 {
	//	t.Fatalf("got client count equal to '%q', expected '1'", length)
	//}
	//
	//chat.Leave(chat.clients[conn1.RemoteAddr().String()])
	//length = len(chat.clients)
	//if length != 0 {
	//	t.Fatalf("got client count equal to '%q', expected '0'", length)
	//}
	//
	//listener.Close()
	//
	////
	////conn2, err := listener.Accept()
	////if err != nil {
	////	t.Fatal(err)
	////}
	////defer conn2.Close()
	////
	////chat.Join(conn2)
	////length = len(chat.clients)
	////if length != 2 {
	////	t.Fatalf("got client count equal to '%q', expected '2'", length)
	////}
}