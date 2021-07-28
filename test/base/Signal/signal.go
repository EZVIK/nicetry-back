package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	signals := make(chan os.Signal, 1)

	done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sg := <-signals
		fmt.Println(sg)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exitting")

}
