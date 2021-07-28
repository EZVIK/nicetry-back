package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
//
//
//
//示例1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//示例 4:
//
//输入: s = ""
//输出: 0
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s由英文字母、数字、符号和空格组成

// abcabcbb
// a = 2 , b = 2, c = 2 , b = 2

func lengthOfLongestSubstring(s string) int {

	subStr := strings.Split(s, "")
	if len(subStr) == 1 {
		return 1
	}

	if len(subStr) == 0 {
		return 0
	}

	subStrTimes := map[string]int{}

	result := map[string]int{}

	start := 0
	end := 0
	i := 0
	for {
		key := subStr[i]

		if subStrTimes[key] == 0 {
			subStrTimes[key]++
			end++
		} else if subStrTimes[key] != 0 {
			subkey := strings.Join(subStr[start:end], "")
			result[subkey] = len(subStrTimes)
			subStrTimes = make(map[string]int)
			i = start
			start++
			end = start
		}

		if end == len(subStr) {
			current_sub := strings.Join(subStr[start:end], "")
			result[current_sub] = len(subStrTimes)
			break
		}

		i++
	}

	max := 0
	sss := ""
	for i, k := range result {
		if k > max {
			max = k
			sss = i
		}
	}

	fmt.Println(sss)
	return max
}

type Point struct {
	Indexs   []int
	Times    int
	IsRepeat bool
}

func main() {

	char := "anviaj"

	fmt.Println(lengthOfLongestSubstring3(char))

}

//
// a n v i a j
// 0 1 2 3 4 5

func lengthOfLongestSubstring3(s string) int {
	r, l := 0, 0

	var ret int

	for i := range s {

		p1 := s[l:i]
		p2 := string(s[i])
		index := strings.Index(p1, p2)

		if index == -1 {
			r++
		} else {
			r = i + 1
			l += index + 1
		}

		p3 := len(s[l:r])
		ret = max(p3, ret)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func test(s string) int {
	r, l := 0, 0

	ret := 0

	for i := range s {
		p1 := s[l:i]
		p2 := string(s[i])
		index := strings.Index(p1, p2)
		if index == -1 {
			r++
		} else {
			r = i + 1
			l += index + 1
		}

		ret = max(len(s[l:r]), ret)
	}

	return ret
}
