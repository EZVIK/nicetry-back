package main

//
//import (
//	"fmt"
//	"os"
//	"strings"
//)
//
//
///*
//	有效数字（按顺序）可以分成以下几个部分：
//
//	 一个 小数 或者 整数
//		（可选）一个 'e' 或 'E' ，后面跟着一个 整数
//		小数（按顺序）可以分成以下几个部分：
//
//	（可选）一个符号字符（'+' 或 '-'）
//		下述格式之一：
//		至少一位数字，后面跟着一个点 '.'
//		至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
//		一个点 '.' ，后面跟着至少一位数字
//
//    整数（按顺序）可以分成以下几个部分：
//	（可选）一个符号字符（'+' 或 '-'）
//		至少一位数字
//
//	部分有效数字列举如下：
//		["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
//
//	部分无效数字列举如下：
//		["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
//	给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。
//*/
//
//
//func main() {
//
//	//numbers := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
//	//numbers := []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"}
//
//	//numbers := []string{".", "+e", "+3. e04116"}
//	//numbers := []string{".1","-.1","1."," 005047e+6","46.e3",".2e81", "4e+"}
//
//	numbers := []string{".1", "1e.", "-1."}
//
//	for _, k := range numbers {
//		fmt.Println(isNumber(k))
//	}
//}
//
//
//func isNumber(number string) bool {
//	number = strings.ToLower(strings.TrimSpace(number))
//	//			  +   -   .   e    E   0  1    2   3   4   5   6   7   8   9
//	d := []int32{43, 45, 46, 101, 69, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
//
//	dicts := make(map[int32]int32)
//
//	for _, value := range d {
//		dicts[value] = value
//	}
//
//	allNumber := true					// without str
//	allString := true
//
//	isMutiStr := make(map[int32][]int32{}, len(number))		// 出现次数不超过2次
//
//	maxStrNum := int32(2)			//
//
//	for FileOp, k := range number {
//
//		if ifLegal := dicts[k]; ifLegal == 0 {
//			return false
//		}
//
//		// no number
//		if k < 48 || k > 57 {
//			allNumber = false
//			isMutiStr[k][0]++
//			isMutiStr[k][1] = int32(FileOp)
//
//			if isMutiStr[k][0] == maxStrNum {
//				return false
//			}
//
//			if len(isMutiStr) > 3 {
//				return false
//			}
//		} else {
//			// number
//			allString = false
//		}
//	}
//
//	// 全字符
//	if allString {
//		return false
//	}
//
//	// 全数字
//	if allNumber {
//		return true
//	}
//
//	if allNumber && len(isMutiStr) > 0 {
//		return false
//	}
//
//	if strings.Contains(number, ".") {
//
//	}
//
//	if strings.Contains(number, "-") {
//
//		if isMutiStr[45][1] == int32(len(number) - 1) {
//			return false
//		}
//	}
//
//	if strings.Contains(number, "+") {
//
//	}
//
//	if strings.Contains(number, "e") {
//
//	}
//
//	return false
//}
