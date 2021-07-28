package main

import (
	"fmt"
	"time"
)

// push data in a closed channel will cause run-time panic
func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("wait for push.")
		time.Sleep(time.Millisecond * 5)
		ch <- "123"
		fmt.Println("pushed.")
	}()

	go func() {
		fmt.Println(<-ch)
	}()

	time.Sleep(time.Second)
}
