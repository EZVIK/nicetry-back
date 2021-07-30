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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
		listA 中节点数目为 m
		listB 中节点数目为 n
		0 <= m, n <= 3 * 104
		1 <= Node.val <= 105
		0 <= skipA <= m
		0 <= skipB <= n
		如果 listA 和 listB 没有交点，intersectVal 为 0
		如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1]
*/
// 	如果两个链表没有交点，返回 null.
//	在返回结果后，两个链表仍须保持原有的结构。
//	可假定整个链表结构中没有循环。
//	程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

// 1 2 3 4
//			8 9 10
//   5 6 7
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	la, lb := 0, 0
	// 1. 遍历查询两个链表的长度 la, lb
	for {
		// 末尾
		if a.Next == nil && b.Next == nil {

			// 2. 判断链表尾部是否相同，不同则说明不🍌
			if a != b {
				return nil
			}
			break
		}

		if a == b {
			return a
		}

		if a.Next != nil {
			a = a.Next
			la++
		}

		if b.Next != nil {
			b = b.Next
			lb++
		}
	}

	// 判断链表长度并对齐长度
	gap := la - lb
	a, b = headA, headB

	if gap > 0 {
		for i := 0; i < gap; i++ {
			a = a.Next
		}
	} else if gap < 0 {
		for i := 0; i < -gap; i++ {
			b = b.Next
		}
	}

	// 判断共同交点
	for {
		if a == b {
			return a
		}

		a = a.Next
		b = b.Next

		if a == nil && b == nil {
			return nil
		}
	}
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
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
	k3 := &ListNode{Val: 5, Next: nil}
	k2 := &ListNode{Val: 4, Next: nil}
	k1 := &ListNode{Val: 8, Next: k2}

	l3 := &ListNode{Val: 1, Next: k3}
	l2 := &ListNode{Val: 6, Next: l3}
	l1 := &ListNode{Val: 5, Next: l2}

	r2 := &ListNode{Val: 1, Next: k1}
	r1 := &ListNode{Val: 4, Next: r2}

	p := getIntersectionNode(r1, l1)

	fmt.Println(p.Val)

	//h := reverseList(header)
	//fmt.Println("asdads", h)

}
