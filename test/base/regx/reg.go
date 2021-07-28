package main

import (
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// r'[+-]?(\d+)?\.?(?(1)\d*|\d+)([eE][+-]?\d+)?
	regexp.Compile("[+-]?(\\d+)?\\.?(?(1)\\d*|\\d+)([eE][+-]?\\d+)?")
}
