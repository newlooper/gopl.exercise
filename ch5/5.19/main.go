package main

import "fmt"

func main() {
	fmt.Println(returnNoZero())
}

func returnNoZero() (result int) {
	defer func() {
		result = recover().(int)
	}()
	panic(4004)
}
