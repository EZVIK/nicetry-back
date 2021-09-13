package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	defer func() {
		fmt.Println("goroutine:", runtime.NumGoroutine())
	}()

	//ch := make(chan int, 1)
	var ch chan int

	// chan 没有进行初始化  读和写都会造成阻塞
	go func() {
		<-ch
	}()

	time.Sleep(time.Second)
}
