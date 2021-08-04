package main

import (
	"fmt"
	"sync"
)

type Item interface{}

type Shit struct {
}

func (s *Shit) String() string {
	return "Your peace of shit"
}

type Stack struct {
	ID            int64
	Capacity      int64
	CurrentLength int64
	Items         []Item
	TopPointer    *Item
	lock          sync.RWMutex
}

func NewStack() *Stack {
	stack := &Stack{}
	stack.Items = []Item{}
	return stack
}

// Print all the elements
func (s *Stack) Print() {
	fmt.Println(s.Items)
}

//
func (s *Stack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.Items) == 0 {
		return nil
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[0 : len(s.Items)-1]
	return item
}

func (s *Stack) Push(item Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Items = append(s.Items, item)
}

//func main() {
//
//	stack := NewStack()
//	stack.Print()
//
//	msgQueue := make(chan string)
//
//	go func() {
//		msgQueue <- "Hi your are here to poop, right?"
//	}()
//
//
//	msg := <- msgQueue
//
//	fmt.Println(msg)
//}

//func main()  {
//	var wg sync.WaitGroup
//	var lock sync.Mutex
//	num := 10
//	n := 100
//	cur := 1
//
//	print := func() {
//		defer wg.Done()
//
//		for {
//			lock.Lock()
//
//			if cur > n {
//				lock.Unlock()
//				break
//			}
//
//			fmt.Printf("%d\n", cur)
//			cur++
//			lock.Unlock()
//		}
//	}
//
//	for FileOp:=0;FileOp<num;FileOp++ {
//		wg.Add(1)
//		go print()
//	}
//
//	wg.Wait()
//}

// 按顺序打印 Cat Dog Fish 100遍
