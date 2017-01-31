package main

import (
	"testing"
	"net"
	"time"
)

func TestChat_Join_Leave(t *testing.T) {
	chat := NewChat()

	length := len(chat.clients)
	if length != 0 {
		t.Errorf("got client count equal to '%q', expected '0'", length)
	}


	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		t.Error(err)
	}

	dial1, err := net.Dial("tcp", ":6666")
	if err != nil {
		t.Error(err)
	}

	dial2, err := net.Dial("tcp", ":6666")
	if err != nil {
		t.Error(err)
	}

	conn1, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}

	chat.Join(conn1)
	length = len(chat.clients)
	if length != 1 {
		t.Errorf("got client count equal to '%q', expected '1'", length)
	}

	conn2, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}

	chat.Join(conn2)
	length = len(chat.clients)
	if length != 2 {
		t.Errorf("got client count equal to '%q', expected '2'", length)
	}


	dial1.Close()
	time.Sleep(time.Microsecond * 100)
	length = len(chat.clients)
	if length != 1 {
		t.Errorf("got client count equal to '%q', expected '1'", length)
	}

	dial2.Close()
	time.Sleep(time.Microsecond * 100)
	length = len(chat.clients)
	if length != 0 {
		t.Errorf("got client count equal to '%q', expected '0'", length)
	}

	conn1.Close()
	conn2.Close()
	listener.Close()
}

