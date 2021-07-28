package main

import (
	"log"
)

//func main() {
//
//
//	total := 0
//
//	m := &sync.Mutex{}
//
//	go func() {
//		fmt.Println(total)
//	}()
//
//	go func() {
//		fmt.Println(total)
//	}()
//
//	for {
//		m.Lock()
//		total++
//		time.Sleep( time.Second)
//		m.Unlock()
//	}
//
//}

type MyErr struct {
	Msg string
}

func main() {
	var e error
	e = GetErr()
	log.Println(e == nil)
}

func GetErr() *MyErr {
	return nil
}

func (m *MyErr) Error() string {
	return "脑子进煎鱼了"
}
