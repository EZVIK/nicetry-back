package main

import (
	"strconv"
)

func main() {

	ch1 := make(chan string, 10)

	for i := 0; i < 10; i++ {
		ch1 <- "HI" + strconv.Itoa(i)
	}

}
