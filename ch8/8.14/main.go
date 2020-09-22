package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

const (
	chatTimeout = 60 * time.Second
	nameTimeout = 10 * time.Second
)

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
				if c != cli {
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

	namePrompt := make(chan string, 2)
	go clientWriter(conn, namePrompt)

	var who string
	input := bufio.NewScanner(conn) // ignoring potential errors from input.Err()

	nameTimer := time.NewTimer(nameTimeout)
	namePrompt <- "Enter your name:"
	select {
	case <-nameTimer.C:
		conn.Close()
		return
	default:
		input.Scan()
		name := strings.Trim(input.Text(), "　 \t\r\n")
		if name == "" {
			who = conn.RemoteAddr().String()
		} else {
			who = name
		}
	}

	namePrompt <- "Welcome, " + who
	close(namePrompt)

	var c client
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	c.outgoing = ch
	c.name = who

	//ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- c

	timer := time.NewTimer(chatTimeout)
	go func() {
		<-timer.C
		log.Printf("client [%s](%s) timeout", who, conn.RemoteAddr())
		conn.Close()
	}()

	for input.Scan() {
		timer.Reset(chatTimeout)
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
