package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 2 * time.Second

func main() {
	//// 创建一个监听8000端口的服务器
	//http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	ctx := r.Context()
	//	// 输出到STDOUT展示处理已经开始
	//	fmt.Fprint(os.Stdout, "processing request\n")
	//	// 通过select监听多个channel
	//
	//	select {
	//	case <-time.After(2 * time.Second):
	//		// 如果两秒后接受到了一个消息后，意味请求已经处理完成
	//		// 我们写入"request processed"作为响应
	//		w.Write([]byte("request processed"))
	//	case <-ctx.Done():
	//
	//		// 如果处理完成前取消了，在STDERR中记录请求被取消的消息
	//		fmt.Fprint(os.Stderr, "request cancelled\n")
	//	}
	//}))

	// creat a context with timeout 1 Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)

	defer cancel()
	//cancel()
	flag := true

	for {

		if !flag {
			break
		}

		select {

		// after 1 second
		case <-time.After(1 * time.Second):
			fmt.Println("FUCK OFF.")

		// timeout
		case <-ctx.Done():
			fmt.Println("Done.", ctx.Err())
			flag = false
		}

	}
}
