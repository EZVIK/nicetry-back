package main

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

*/

// ((()){{}{}[]})
// {}()[]

// (()

/**
1 loop 将字符 依次放入堆中
2 last init
3 放入字符 并与 last 比较
	4 如果同向 则继续
	5 如果是异向
		6 同类型 直接删除
		7 不同类型直接 return false
*/

//var start, end = 0, 1

func main() {

	// 40 41, 123 125, 91 93
	//isValid("(){}[]")
	//fmt.Println(isValid3("({{}})"))
	//fmt.Println(f(106))
	fmt.Println(fuck(6))
}

var CostRate = []int{5, 20, 50, 100, 500, 1000, 2000, 3000, 4000, 5000, 6000, 6001}
var CostWeight = []int{30, 15, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

func fuck(a int) int {
	sum := 0
	index := 0
	for i := len(CostRate) - 1; i >= 0; i-- {
		if a >= CostRate[i] {
			index = i
			break
		}
	}
	sum += ans(index)
	sum += (a - CostRate[index]) * CostWeight[index+1]
	return sum
}

func ans(index int) int {
	if index == 0 {
		return CostRate[index] * CostWeight[index]
	}
	return ans(index-1) + CostRate[index]*CostWeight[index]
}

func isValid3(s string) bool {
	if len(s) == 0 {
		return true
	}

	brackets := map[byte]byte{')': '(', ']': '[', '}': '{'}
	checkArr := make([]byte, 0, len(s))

	for _, v := range s {
		if _, ok := brackets[byte(v)]; ok {

			if brackets[byte(v)] != checkArr[len(checkArr)-1] {
				return false
			}

			checkArr = checkArr[:len(checkArr)-1]
		} else {
			checkArr = append(checkArr, byte(v))
		}
	}

	if len(checkArr) != 0 {
		return false
	}

	return true

}

func isValid2(s string) bool {

	if s == "" {
		return true
	}

	checkMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	checkStack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			checkStack = append(checkStack, s[i])
		} else {
			if checkStack[len(checkStack)-1] != checkMap[s[i]] {
				return false
			}
			checkStack = checkStack[:len(checkStack)-1]
		}
	}

	return len(checkStack) == 0

}

func isValid(s string) bool {

	arr := []byte(s)

	last := arr[0]

	st := stack()

	for i := range arr {

		if len(st.ElementArr) == 0 {
			st.push(arr[i])
			last = arr[i]
			continue
		}

		ppp := arr[i]
		temp_res := check(last, ppp)

		if temp_res == 1 {
			st.push(ppp)
			last = ppp
		} else if temp_res == 0 {
			st.pop()

			if len(st.ElementArr) == 0 {
				continue
			}

			last = st.ElementArr[len(st.ElementArr)-1]
		} else {
			return false
		}
	}

	if len(st.ElementArr) == 0 {
		return true
	}

	return false
}

func check(last, current byte) int {

	if last == 40 {
		if current == 125 || current == 93 {
			return -1
		} else if current == 41 {
			return 0
		} else {
			return 1
		}
	} else if last == 123 {
		if current == 41 || current == 93 {
			return -1
		} else if current == 125 {
			return 0
		} else {
			return 1
		}
	} else if last == 91 {
		if current == 41 || current == 125 {
			return -1
		} else if current == 93 {
			return 0
		} else {
			return 1
		}
	} else {
		return -1
	}
}

type Stack struct {
	ElementArr []byte
}

func stack() *Stack {
	s := new(Stack)
	s.ElementArr = make([]byte, 0)
	return s
}

func (s *Stack) pop() interface{} {

	if len(s.ElementArr) == 0 {
		return nil
	}
	pop := s.ElementArr[len(s.ElementArr)-1]
	s.ElementArr = s.ElementArr[:len(s.ElementArr)-1]

	return pop
}

func (s *Stack) push(i byte) {
	s.ElementArr = append(s.ElementArr, i)
}
