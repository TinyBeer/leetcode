package main

import "math"

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左
子树
只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var df func(*TreeNode, int, int) bool
	df = func(tn *TreeNode, l, r int) bool {
		if tn.Left != nil {
			if tn.Left.Val >= tn.Val || tn.Left.Val <= l || !df(tn.Left, l, min(r, tn.Val)) {
				return false
			}
		}

		if tn.Right != nil {
			if tn.Right.Val <= tn.Val || tn.Right.Val >= r || !df(tn.Right, max(tn.Val, l), r) {
				return false
			}
		}

		return true
	}
	return df(root, math.MinInt, math.MaxInt)
}
