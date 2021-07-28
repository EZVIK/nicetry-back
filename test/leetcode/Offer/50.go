package main

import "fmt"

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

func firstUniqChar(s string) byte {
	m := make([]int, 26)

	for _, v := range s {
		m[v-97] += 1
	}

	for _, v := range s {
		if m[v-97] == 1 {
			return byte(v)
		}
	}

	return byte(32)
}

func firstUniqChar1(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	s := "abaccdeff"

	//arr := make([]int, len(s))

	fmt.Println(string(firstUniqChar(s)))

}
