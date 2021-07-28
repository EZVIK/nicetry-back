package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	time.Sleep(10)
	waitForSignal2()

}

// 监听 os signal
// 当程序退出时 像 chan sigs 发送数据
func waitForSignal2() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}
