package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

	输入：l1 = [1,2,4], l2 = [1,3,4]
	输出：[1,1,2,3,4,4]


	输入：l1 = [1,2,4], l2 = [2,3,4]
	输出：[1,2,2,3,4,4]

	两个链表的节点数目范围是 [0, 50]
	-100 <= Node.val <= 100
	l1 和 l2 均按 非递减顺序 排列

*/

//type ListNode struct {
//	 Val int
//	 Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	l, j := l1, l2

	head := new(ListNode)

	h := head

	for {

		if l == nil && j == nil {
			return head.Next
		}

		if l == nil {
			h.Next = j
			h = h.Next
			j = j.Next
			return head.Next

		}

		if j == nil {
			h.Next = l
			h = h.Next
			l = l.Next
			return head.Next
		}

		if l.Val > j.Val {
			h.Next = j
			h = h.Next
			j = j.Next
		} else if l.Val < j.Val {
			h.Next = l
			h = h.Next
			l = l.Next
		} else {
			h.Next = j
			h = h.Next
			j = j.Next

			h.Next = l
			h = h.Next
			l = l.Next
		}

	}
}

func getNodeDeep21(l *ListNode) int {
	if l.Next != nil {
		s := getNodeDeep21(l.Next) + 1
		return s
	} else {
		return 1
	}
}

func InitNodeFromArr21(arr []int) *ListNode {
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

func main() {

	arr1 := []int{-8, -7, -4, 0, 7, 8, 9}
	arr2 := []int{10}

	k1 := InitNodeFromArr21(arr1)
	k2 := InitNodeFromArr21(arr2)
	k3 := mergeTwoLists(k1, k2)

	fmt.Println(k3)
}
