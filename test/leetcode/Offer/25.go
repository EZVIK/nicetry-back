package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

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

		if l.Val < j.Val {
			h.Next = l
			h = h.Next
			l = l.Next
		} else if l.Val > j.Val {
			h.Next = j
			h = h.Next
			j = j.Next
		} else {
			h.Next = l
			h = h.Next
			l = l.Next

			h.Next = j
			h = h.Next
			j = j.Next
		}

	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	h := head
	var pre *ListNode
	cur := h
	next := &ListNode{}

	for cur != nil {
		next = cur.Next // next = current.next;
		cur.Next = pre  // current.next = prev;
		pre = cur       // prev = current;
		cur = next      // current = next;
	}

	h = pre
	return h
}

func main() {

	l3 := &ListNode{Val: 4, Next: nil}
	l2 := &ListNode{Val: 3, Next: l3}
	l1 := &ListNode{Val: 2, Next: l2}
	header := &ListNode{Val: 1, Next: l1}

	h := reverseList(header)
	fmt.Println("asdads", h)
}
