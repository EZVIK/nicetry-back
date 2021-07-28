package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:' root ='[3,5,1,6,2,0,8,null,null,7,4]

*/

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func PreOrderTraversal(node *TreeNode1) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PreOrderTraversal(node.Left)
	PreOrderTraversal(node.Right)
}

func main() {

}

//func (this *TreeNode) IfSubNode(node *TreeNode) bool {
//	if this.Left != nil {
//
//	}
//}

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//
//}
