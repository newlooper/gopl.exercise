package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	outgoing chan<- string // outgoing message channel
	name     string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

const timeout = 60 * time.Second

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.outgoing <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			for c := range clients {
				if c == cli {
					cli.outgoing <- "Welcome!"
				} else {
					cli.outgoing <- "Online user: " + c.name
				}
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outgoing)
		}
	}
}

func handleConn(conn net.Conn) {
	var c client
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	c.outgoing = ch
	c.name = who

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- c

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		log.Printf("client [%s] timeout", conn.RemoteAddr())
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timer.Reset(timeout)
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- c
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
