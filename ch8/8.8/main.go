package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const timeout = 10

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {

	defer c.Close()

	input := bufio.NewScanner(c)
	timer := time.NewTimer(timeout * time.Second)
	msgs := make(chan string)
	wg := sync.WaitGroup{}

	go func() {
		for input.Scan() {
			msgs <- input.Text()
		}
		close(msgs)
	}()

	for {
		select {
		case msg, ok := <-msgs:
			if ok {
				wg.Add(1)
				go func() {
					defer wg.Done()
					echo(c, msg, 1*time.Second)
				}()
				timer.Reset(timeout * time.Second)
			} else {
				wg.Wait()
				return
			}
		case <-timer.C:
			timer.Stop()
			fmt.Printf("no message from client [%s] for %d seconds, server drop it.", c.RemoteAddr(), timeout)
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
