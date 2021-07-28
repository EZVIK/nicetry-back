package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
	age  int8
}

var u = user{name: "A", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "C"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

func main() {
	c := make(chan *user, 5)
	c <- g

	fmt.Println(g) // user{name: "Ankur", age: 25}

	// modify g	 g-> &user{name: "Ankur Anand", age: 100}
	g = &user{name: "B", age: 100}

	go printUser(c)  // "printUser goRoutine called"
	go modifyUser(g) //
	time.Sleep(5 * time.Second)
	fmt.Println(g)
}
