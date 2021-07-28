package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/**
   	1、 how is the "for range channel" will be compiled
	2、 what's the different of buffered channel and unbuffered channel in this case
	3、 函数中关于 channel 参数的流线是如何校验的
*/

func main() {
	log.Printf("Starting program...")
	// creating a blocking channel
	// note: in certain cases, you can create a buffered channel
	// so that the sender goroutines can quit early
	results := make(chan string)
	go sender(results)
	receiver(results)
}

var globalCount uint64 = 0

// results chan string 接收ch
func sender(results chan<- string) {

	// 创建 waitgroup 等待goroutine 结束
	wg := &sync.WaitGroup{}

	// 创建 5个 goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()                  // 该 协程 结束时 执行wg.Done()
			p := fmt.Sprintf("result %d", i) // 往 results ch 传入 p

			// if channel is unbuffered , it will wait for util the <- chan
			results <- p

			log.Printf("Pushing %s\n", p)

		}(i) // 方法参数
	}
	// after all the results are sent, we can now
	// safely close the channel
	wg.Wait() // 等待 所有协程结束
	log.Printf("Closing program...")
	close(results) // 关闭 results通道
	log.Printf("Closing program...")
	log.Printf("globalCount:%v", globalCount)

}

// 接收函数
// results 从通道获取 get
func receiver(results <-chan string) {

	// 循环阻塞等待 results 通道， 从通道获取数据直到通道关闭
	for r := range results {
		log.Printf("-- Getting %s\n", r)
		time.Sleep(time.Second) // 阻塞 1s
	}
}
