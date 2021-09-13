package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "123456"

	fmt.Println(strings.Index(s, "6") != len(s)-1)
}
