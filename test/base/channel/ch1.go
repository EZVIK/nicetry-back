package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 编程题：3个函数分别打印cat、dog、fish，
// 要求每个函数都要起一个goroutine，
// 按照cat、dog、fish顺序打印在屏幕上100次。

const totalPrintTime = 2

func dog(wg *sync.WaitGroup, counter uint64, from, to chan struct{}) {
	for {
		if counter >= uint64(totalPrintTime) {
			wg.Done()
			return
		}

		<-from
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		to <- struct{}{}
	}
}

func fish(wg *sync.WaitGroup, counter uint64, from, to chan struct{}) {
	for {
		if counter >= uint64(totalPrintTime) {
			wg.Done()
			return
		}

		<-from
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		to <- struct{}{}
	}
}

func cat(wg *sync.WaitGroup, counter uint64, from, to chan struct{}) {
	for {
		if counter >= uint64(totalPrintTime) {
			wg.Done()
			return
		}

		<-from
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		to <- struct{}{}
	}
}

func Animal(name string, wg *sync.WaitGroup, counter uint64, from, to chan struct{}) {
	for {
		if counter >= uint64(totalPrintTime) {
			wg.Done()
			return
		}

		// 用Channel 控制流向 堵塞
		<-from
		fmt.Println(name)
		atomic.AddUint64(&counter, 1)
		to <- struct{}{}
	}
}

func main() {
	var wg sync.WaitGroup
	var counter uint64
	//var fishcounter uint64
	//var dogcounter uint64

	fishch := make(chan struct{}, 1)
	catch := make(chan struct{}, 1)
	dogch := make(chan struct{}, 1)

	wg.Add(3)
	dogch <- struct{}{}
	go Animal("Dog", &wg, counter, dogch, fishch)
	go Animal("Cat", &wg, counter, fishch, catch)
	go Animal("Fish", &wg, counter, catch, dogch)
	wg.Wait()
}
