package main

import (
	"fmt"
)

// give an header node of a linklist, reverse all link node, return a array

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode1 struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	length := 0
	h := head
	for h != nil {
		h = h.Next
		length++
	}

	result := make([]int, length)

	for i := length - 1; i >= 0; i-- {
		result[i] = head.Val
		head = head.Next
	}

	return result
}

// return reverse header (tail node)
// 1 -> 2 -> 3 -> nil
// 3 -> 2 -> nil
// 2 ->

// input, output
func re(n *ListNode) *ListNode {

	var pre *ListNode
	cur := n
	var next *ListNode

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	n = pre
	return n
}

func main() {

	l3 := &ListNode{Val: 4, Next: nil}
	l2 := &ListNode{Val: 3, Next: l3}
	l1 := &ListNode{Val: 2, Next: l2}
	header := &ListNode{Val: 1, Next: l1}

	//empty := &ListNode{Val: 1, Next: nil}

	res := re(header)
	fmt.Println(res)

}
