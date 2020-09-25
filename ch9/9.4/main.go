package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func init() {
	println("press ENTER to quit...")
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancel()
	}()
}

const CHANNEL_NUMBER = 1000000 // suggestion: do not set a large number on windows

var ctx, cancel = context.WithCancel(context.Background())
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go memoryProbe()

	wg.Add(1)
	go run()

	wg.Wait()
}

func run() {
	defer wg.Done()
	var chs [CHANNEL_NUMBER]chan string
	for i := 0; i < CHANNEL_NUMBER; i++ {
		chs[i] = make(chan string)
		if i < CHANNEL_NUMBER-1 {
			go func(i int) {
				chs[i+1] <- <-chs[i] // before `chs[0] <- "Go"`, all goroutine will block, so needn't worry about "Index out of bounds"
				close(chs[i])
			}(i)
		}
	}

	start := time.Now()

	chs[0] <- "Go"          // first channel
	<-chs[CHANNEL_NUMBER-1] // last channel

	fmt.Printf("\r\n%d goroutines, %fs\n", CHANNEL_NUMBER, time.Since(start).Seconds())
}

func memoryProbe() {
	defer wg.Done()
	var m runtime.MemStats
	for {
		select {
		case <-ctx.Done():
			return
		default:
			runtime.ReadMemStats(&m)
			fmt.Printf("\033[2K\rAlloc=%vm\tTotalAlloc=%vm\tSys=%vm\tNumGC=%v",
				m.Alloc/1024/1024,
				m.TotalAlloc/1024/1024,
				m.Sys/1024/1024,
				m.NumGC)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

/*
go run .
press ENTER to quit...
Alloc=459m      TotalAlloc=473m Sys=2135m       NumGC=6
1000000 goroutines, 0.335577s
Alloc=572m      TotalAlloc=600m Sys=2669m       NumGC=7
*/
