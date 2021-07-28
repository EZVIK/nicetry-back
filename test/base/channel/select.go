package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 1
		}
	}()

	go func() {
		for range time.Tick(2 * time.Second) {
			<-ch
		}
	}()

	for {
		select {
		case <-ch:
			{
				fmt.Println("case1")
			}
		case ch <- 0:
			{
				fmt.Println("case2")
			}
		}
	}
}
