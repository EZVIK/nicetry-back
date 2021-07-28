package main

import "fmt"

func fibonacci(seq int, ch chan int64) {

	x, y := int64(0), int64(1)
	for i := 0; i < seq; i++ {
		ch <- x
		x, y = y, x+y
	}

	close(ch)
}

func main() {

	c := make(chan int64, 33)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

}
