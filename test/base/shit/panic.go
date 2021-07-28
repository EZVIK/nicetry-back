package main

import "fmt"

func main() {
	fmt.Println("start")

	defer func() {
		fmt.Println("defer start")

		if err := recover(); err != nil {
			fmt.Println("recover start")
			fmt.Println(err)
		}

		fmt.Println("defer end")

	}()

	f()

	fmt.Println("end")
}

func f() {
	fmt.Println("添加 f start")

	panic("f uck ")
}
