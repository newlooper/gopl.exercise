package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	ping, pong := make(chan int64), make(chan int64)
	var i int64

	start := time.Now()
	go func() {
		ping <- 1
		for {
			i++
			ping <- <-pong
		}
	}()
	go func() {
		for {
			pong <- <-ping
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Printf("\n%.0f round / s\n",float64(i)/float64(time.Since(start))*1e9)
}

/*
windows 10
go run .
2422471 round / s
*/

/*
ubuntu 20.04
go run .
^C
3738901 round / s
*/
