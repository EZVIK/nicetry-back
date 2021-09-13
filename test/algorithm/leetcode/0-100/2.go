package main

import (
	"fmt"
)

//type ListNode struct {
//	Val int
//  	Next *ListNode
//}

func GetEndNode(l *ListNode) *ListNode {
	if l.Next == nil {
		return l
	} else {
		return GetEndNode(l.Next)
	}
}

func AdddNilNodeFromEnd(times int, ll *ListNode) *ListNode {
	times -= 1
	if times == 0 {
		l := new(ListNode)
		ll.Next = l
		return l
	} else {
		last := AdddNilNodeFromEnd(times, ll)
		l := new(ListNode)
		last.Next = l
		return l
	}
}

func getNodeDeep(l *ListNode) int {
	if l.Next != nil {
		s := getNodeDeep(l.Next) + 1
		return s
	} else {
		return 1
	}
}

func InitNodeFromArr(arr []int) *ListNode {
	head := new(ListNode)
	first := head
	for i, k := range arr {
		l2 := new(ListNode)

		l := new(ListNode)
		l.Val = k

		if i != len(arr)-1 {
			l.Next = l2
		}

		first.Next = l

		first = l
	}
	return head.Next
}

func addTwoNumbersNew(l1 *ListNode, l2 *ListNode) *ListNode {

	head := new(ListNode)

	h := head

	carry := 0

	for {

		if l1 == nil && l2 == nil {
			if carry > 0 {
				last := new(ListNode)
				last.Val = 1
				h.Next = last
			}
			return head.Next
		}

		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}

		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		carry = 0
		temp_sum := sum % 10

		if sum >= 10 {
			carry++
		}

		h.Next = &ListNode{Val: temp_sum}

		h = h.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{Val: 0, Next: nil}
	current := head
	carry := 0

	for l1 != nil || l2 != nil {

		var x, y int

		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		current.Next = &ListNode{Val: (x + y + carry) % 10, Next: nil}

		current = current.Next

		carry = (x + y + carry) / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry % 10, Next: nil}
	}

	return head.Next
}

func t(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil || l2 == nil {
		return nil
	}

	head := new(ListNode)
	current := head
	carry := 0
	for {

		if l1 == nil && l2 == nil {
			break
		}

		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}

		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		current.Next = &ListNode{Val: sum % 10, Next: nil}
		current = current.Next

		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry, Next: nil}
	}

	return head.Next
}

func main() {

	//arr1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	//arr2 := []int{5,6,4}

	arr1 := []int{9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9, 9, 9, 9}

	k1 := InitNodeFromArr(arr1)
	k2 := InitNodeFromArr(arr2)

	head := addTwoNumbersNew(k2, k1)

	fmt.Println("k1", k2, head)

	//fmt.Println(math.Log10(564))
	//fmt.Println(int(math.Log10(1000000000000000000000000000001)))
}

//You are given two non-empty linked lists representing two non-negative integers.
//
//The digits are stored in reverse order, and each of their nodes contains a single digit.
//
//Add the two numbers and return the sum as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.

//输入：l1 = [2,4,3], l2 = [1,0,0,0,0,5,6,4]
//输出：[7,10k,8]
//解释：342 + 46500001 = 807.
