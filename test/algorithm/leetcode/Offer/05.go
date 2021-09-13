package main

import (
	"fmt"
	"strings"
)

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
func replaceSpace(s string) string {
	builder := strings.Builder{}
	builder.Grow(len(s))
	last := 0

	for i := range s {
		if s[i] == ' ' {
			builder.WriteString(s[last:i])
			builder.WriteString("%20")
			last = i + 1
		}
	}

	builder.WriteString(s[last:])
	return builder.String()
}

func replaceSpace2(s string) string {
	ans := ""

	last := 0
	for i, v := range s {
		if v == 32 {
			ans += s[last:i]
			ans += "%20"
			last = i + 1
		}
	}
	ans += s[last:]
	return ans
}

func main() {

	s := "asi 0s"

	fmt.Println(replaceSpace(s))
}
