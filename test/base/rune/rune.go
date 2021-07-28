package main

import (
	"fmt"
	"strconv"
)

func TestOther() {
	text := "abcd1234浮生无事"
	fmt.Println([]byte(text))
	fmt.Println([]rune(text))
}

func TestOther7() {
	text := "abcd1234浮生无事"
	textLen := len(text)
	fmt.Println("len:" + strconv.FormatInt(int64(textLen), 10))
	for i := 0; i <= textLen-1; i++ {
		fmt.Printf("word:%s\n", text[i:i+1])
	}
}

func TestOther6() {
	text := "abcd1234浮生无事"
	textRune := []rune(text)
	textLen := len(textRune)

	fmt.Println("len:" + strconv.FormatInt(int64(textLen), 10))

	for i := 0; i <= textLen-1; i++ {
		fmt.Printf(fmt.Sprintf("word:%s\n", string(textRune[i:i+1])))
	}
}

func main() {

	TestOther7()
	fmt.Println()
	TestOther6()
}
