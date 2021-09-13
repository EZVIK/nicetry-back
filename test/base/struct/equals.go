package main

import "fmt"

type Customer struct {
	ID   int
	Name string
	Arr  []tt
}

type tt struct {
	ID int
}

func main() {
	a := new(struct{})
	b := new(struct{})
	c := struct{}{}
	println(a, b, &c, a == b, b == &c)
	//println(*b == c)
	fmt.Println(a, b)
	println(a, b, &c, a == b, b == &c)

}
