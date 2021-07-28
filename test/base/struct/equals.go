package main

import "fmt"

type Customer struct {
	ID   int
	Name string
	Arr  []int
}

func main() {
	a := []int{}
	var c1 interface{} = Customer{ID: 1, Name: "Jack", Arr: a}
	var c2 interface{} = Customer{ID: 1, Name: "Jack", Arr: a}
	fmt.Println(c1 == c2)

}
