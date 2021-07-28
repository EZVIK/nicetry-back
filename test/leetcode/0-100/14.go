package main

import (
	"fmt"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串""。


示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成

*/
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	min_len := len(strs[0])
	for _, k := range strs {
		l := len(k)

		if l == 0 {
			return ""
		} else {
			if l < min_len {
				min_len = l
			}
		}

	}
	start := 0

	// loop 1 ,
	for {

		flag := true

		temp_prex := strs[0][start]

		for i := 1; i < len(strs); i++ {
			k := strs[i]

			if k[start] != temp_prex {
				flag = false
				break
			}

			if start >= len(k) {
				return strs[0][:start+1]
			}
		}

		if !flag {
			return strs[0][:start]
		}

		start++

		if start >= min_len {
			return strs[0][:start]
		}
	}
}

func main() {

	params := []string{"axfla", "aflaow", "aflaiaght"}
	d := longestCommonPrefix(params)
	fmt.Println(">>>", d, ">>>")
}
