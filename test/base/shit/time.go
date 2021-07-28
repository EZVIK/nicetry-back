package main

import (
	"fmt"
	"sync"
	"time"
)

func Cat() string {
	return "1 Cat"
}

func Dog() string {
	return "2 Dog"
}

func Fish() string {
	return "3 Fish"
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Print("worker starting: ", id)

	time.Sleep(time.Millisecond)

	fmt.Print("worker end: ", id)

}

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("END")
}
