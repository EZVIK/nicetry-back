package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

/**
调用第三方请求时 请求的 timeout 设置的过长 或者 不设置, 可能会由于协程等待请求结果造成阻塞，
*/
func main() {
	fmt.Println("Starting: ", runtime.NumGoroutine())
	httpC := http.Client{
		Timeout: time.Second * 1,
	}
	for {
		ans := runtime.NumGoroutine()
		fmt.Println("For: ", ans)

		if ans < 30 {
			go func() {
				fmt.Println("Go: ", runtime.NumGoroutine())

				_, err := httpC.Get("https://www.xxx.com/")

				if err != nil {
					fmt.Printf("http.Get err: %v\n", err)
				}
			}()
		}

		time.Sleep(time.Second * 1)

		//fmt.Println("goroutines: ", runtime.NumGoroutine())
	}
}
