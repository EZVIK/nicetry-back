package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan string, 1)

	go func() {
		a <- "123"
	}()

	ha := a

	go func() {
		fmt.Println(ha, len(ha), cap(ha))
		fmt.Println(a, len(a), cap(a))

		fmt.Println(<-ha)

		fmt.Println(a, len(a), cap(a))
	}()

	time.Sleep(time.Millisecond)
}
