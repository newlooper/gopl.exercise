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

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go func() { // 若不想用匿名函数，则要么 wg 设为包级变量，要么 wg 作为参数传给 echo，这样 defer wg.Done() 就可以写在 echo 函数体内
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}
	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	if c, ok := c.(*net.TCPConn); ok { // 断言，放大方法集
		c.CloseWrite()
	} else {
		c.Close()
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
