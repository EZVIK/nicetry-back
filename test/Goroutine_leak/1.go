package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

/**
 * @Description:

	channel 使用不当导致出现 Goroutine 泄漏
*/
func main() {

	for {
		queryAll()
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}

func queryAll() int {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i, "start.")
			ch <- query()
			fmt.Println(i, "end.")
		}(i)
	}

	return <-ch
}

func query() int {
	n := rand.Intn(100)
	//time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}
