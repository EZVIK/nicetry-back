package main

import (
	"fmt"
	"sync"
	"time"
)

func Out(index int, ch chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 3)
	fmt.Println("[", index, "]", <-ch)
	wg.Done()
}

func main() {

	size, n := 100, 0
	ch := make(chan int, size)
	var wg sync.WaitGroup

	for n < size {
		n++
		ch <- n
	}

	times := 0

	for times < 100 {
		wg.Add(1)
		go Out(times, ch, &wg)
		times++
	}

	wg.Wait()

	fmt.Println("end")
}
