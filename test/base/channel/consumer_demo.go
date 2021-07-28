package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	stop := make(chan bool)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int, stop chan bool) {
			defer wg.Done()

			consumer(index, stop)
		}(i, stop)

	}

	waitForSignal()
	close(stop)
	fmt.Println("stopping all jobs!")
	wg.Wait()
}

// 消费协程
func consumer(id int, stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Printf("%v exit sub goroutine.\n", id)
			return
		default:
			fmt.Println("running...", id)
			time.Sleep(time.Second)
		}
	}
}

// 监听 os signal
// 当程序退出时 像 chan sigs 发送数据
func waitForSignal() {
	fmt.Println("waitForSignal......")
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	fmt.Println("Notify......")

	<-sigs
}
