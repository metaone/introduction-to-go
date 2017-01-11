package main

import (
	"flag"
	"net"
	"log"
	"io"
)

func main() {
	serverPath := flag.String("server", "localhost:11114", "host:port")
	proxyPath := flag.String("proxy", "localhost:11115", "host:port")
	flag.Parse()

	server, err := net.Listen("tcp", *serverPath)
	if err != nil {
		log.Panic(err)
	}
	defer server.Close()

	proxy, err := net.Listen("tcp", *proxyPath)
	if err != nil {
		log.Panic(err)
	}
	defer proxy.Close()


	redirect, err := net.Dial("tcp", *serverPath)
	if err != nil {
		panic(err)
	}
	defer redirect.Close()

	for {
		connServer, err := server.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handle(connServer, connServer)

		connProxy, err := proxy.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handle(connProxy, connProxy)
		go handle(connProxy, redirect)
	}
}

func handle(conn net.Conn, redirect net.Conn)  {
	_, err := io.Copy(conn, redirect)
	if err != nil {
		log.Panic(err)
	}
	conn.Close()
}

