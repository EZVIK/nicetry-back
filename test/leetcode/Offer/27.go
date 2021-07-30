package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

 * 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
 */

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func mirrorTree(root *TreeNode) *TreeNode {
	n := root
	if n == nil {
		return nil
	}

	if n.Left == nil && n.Right == nil {
		return n
	} else if n.Left != nil && n.Right != nil {
		temp := n.Left
		n.Left = n.Right
		n.Right = temp
	} else if n.Left == nil && n.Right != nil {
		n.Left = n.Right
		n.Right = nil
	} else if n.Left != nil && n.Right == nil {
		n.Right = n.Left
		n.Left = nil
	} else {
		return nil
	}

	mirrorTree(n.Left)
	mirrorTree(n.Right)

	return n
}

func main() {

	//l6 := &TreeNode{Val: 9, Left: nil, Right: nil}
	//l5 := &TreeNode{Val: 6, Left: nil, Right: nil}
	//l4 := &TreeNode{Val: 3, Left: nil, Right: nil}
	//l3 := &TreeNode{Val: 1, Left: nil, Right: nil}
	//l2 := &TreeNode{Val: 7, Left: l5, Right: nil }
	//l1 := &TreeNode{Val: 2, Left: l3, Right:l4}
	//root := &TreeNode{Val: 4, Left: l1, Right: l2}

	//n := root
	//mirrorTree(l6)
	//
	//fmt.Println("SSD")
}
