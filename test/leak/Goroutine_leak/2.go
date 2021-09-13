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

	// 创建了一个 struct{} 类型的 无缓存 channel
	var ch chan struct{}

	// 使用goroutine 调用一个匿名方法 往 这个 channel 入一个 struct{}
	go func() {

		fmt.Println("尝试发送")
		ch <- struct{}{}
		// 发送完毕 但是因为无channel缓存，且没有接收的动作，程序会阻塞在发送的位置，直到睡眠结束，主线程结束并调用defer 方法
		// 两个goroutine 应该分别是 主线程 以及 此方法的协程。

	}()

	// 睡眠
	time.Sleep(time.Second)
}
