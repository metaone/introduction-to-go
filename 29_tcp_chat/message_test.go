package main

import "testing"

func TestMessage(t *testing.T)  {
	msg := &Message{}

	if msg.fromId != "" {
		t.Fatalf("got fromId equals to '%q', expected ''", msg.fromId)
	}

	if msg.from != "" {
		t.Fatalf("got from equals to '%q', expected ''", msg.from)
	}

	if msg.text != "" {
		t.Fatalf("got text equals to '%q', expected ''", msg.text)
	}


	msg = &Message{fromId: "foo", from: "bar", text: "xyz"}

	if msg.fromId != "foo" {
		t.Fatalf("got fromId equals to '%q', expected '%q'", msg.fromId, "xyz")
	}

	if msg.from != "bar" {
		t.Fatalf("got from equals to '%q', expected '%q'", msg.from, "bar")
	}

	if msg.text != "xyz" {
		t.Fatalf("got text equals to '%q', expected '%q'", msg.text, "xyz")
	}
}
