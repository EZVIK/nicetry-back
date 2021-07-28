package main

import "fmt"

func Three_Component() {
	sum := 0

	for i := 0; i < 100; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func WhileLoop() {

	for n := 0; n < 5; {
		fmt.Println(n)
		n++
	}

	n := 0
	for n < 100 {
		fmt.Println(n)
		n++
	}

}

func Infinite() {
	for {
		fmt.Println("SHITING")
	}
}

func For_Each() {
	strings := []string{"HELLO", "BITCH"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}

func Exit_loop() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 {
			continue
		}
		sum += 1
	}
}

func main() {

	a := uint(1)
	b := uint(2)
	// 1 - 2
	// 0 - 1
	// 0 + (-1) 补码
	// uint without operator
	//
	// 00000 + 111111
	fmt.Println(a - b)

	cc := []rune{'a', 'v', 'b', 'w'}

	for k, v := range cc {
		fmt.Println(k, v)
	}
}
