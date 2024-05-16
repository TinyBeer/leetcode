package main

import (
	"strconv"
	"strings"
)

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) string {
	var res []string
	var df func(*TreeNode)
	df = func(tn *TreeNode) {
		if tn == nil {
			res = append(res, "null")
			return
		}
		res = append(res, strconv.Itoa(tn.Val))
		df(tn.Left)
		df(tn.Right)
	}
	df(root)
	return "[" + strings.Join(res, ",") + "]"
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	_flatten(root)
}

func _flatten(root *TreeNode) (tail *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root
	}

	var curTail *TreeNode = root
	var tempRight = root.Right
	if root.Left != nil {
		curTail = _flatten(root.Left)
		root.Right = root.Left
	}

	if tempRight != nil {
		curTail.Right = tempRight
		curTail = _flatten(tempRight)
	}

	root.Left = nil
	curTail.Left = nil
	return curTail
}

// func flatten(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	var df func(*TreeNode) (*TreeNode, *TreeNode)
// 	df = func(tn *TreeNode) (root *TreeNode, leave *TreeNode) {
// 		root = tn
// 		leave = tn
// 		if tn.Right != nil {
// 			tn.Right, leave = df(tn.Right)
// 		}
// 		if tn.Left != nil {
// 			lr, ll := df(tn.Left)
// 			ll.Right = tn.Right
// 			tn.Right = lr
// 			tn.Left = nil
// 			if leave == tn {
// 				leave = ll
// 			}
// 		}
// 		return
// 	}

// 	df(root)
// }
