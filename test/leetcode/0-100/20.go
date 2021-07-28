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

	// 						{
	fmt.Println(isValid("({){})"))
	//fmt.Println(isValid("(){}[]"))

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
			if len(checkStack) == 0 {
				return false
			}
			if checkStack[len(checkStack)-1] != checkMap[s[i]] {
				return false
			} //if>>>
			checkStack = checkStack[:len(checkStack)-1]
		} //else>>
	} // for>

	return len(checkStack) == 0

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
