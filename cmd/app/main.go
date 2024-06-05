package main

import (
	"log"
	"net"
)

const (
	port     = "8080"
	safeMode = false
)

func safeRemoteAddr(conn net.Conn) string {
	if safeMode {
		return "[REDACTED]"
	}

	return conn.RemoteAddr().String()
}

func handleConnection(conn net.Conn) {
	log.Println("Connected by", safeRemoteAddr(conn))

	message := "Hello, Cosmos!\n"

	n, err := conn.Write([]byte(message))
	if err != nil {
		log.Printf("Unable to write message to %s: %s\n", safeRemoteAddr(conn), err)
		return
	}
	defer conn.Close()

	if n < len(message) {
		log.Printf("Message not fully written to client: %d/%d\n", n, len(message))
		return
	}
}

func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Unable to listen on port %s: %s\n", port, err)
	}

	log.Printf("TCP Server started, listening on port %s...\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Unable to accept connection:\n", err)
			continue
		}

		go handleConnection(conn)
	}
}
