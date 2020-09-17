package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func watch(w io.Writer, r io.Reader, city string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%-10s %s\n", city, s.Text())
	}
	if s.Err() != nil {
		log.Printf("error from %s: %s\n", city, s.Err())
	}
}

func askTime(cp []string) {
	conn, err := net.Dial("tcp", cp[1])
	if err != nil {
		log.Fatal(err)
	}
	watch(os.Stdout, conn, cp[0])
}

// NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
func main() {
	for _, city := range os.Args[2:] {
		cp := strings.Split(city, "=")
		go askTime(cp)
	}

	cp := strings.Split(os.Args[1], "=")
	askTime(cp)
}
