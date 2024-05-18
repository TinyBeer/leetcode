package main

import "math"

/*
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MinInt

	var df func(*TreeNode) int
	df = func(tn *TreeNode) int {
		val := tn.Val
		l, r := 0, 0
		if tn.Left != nil {
			l = df(tn.Left)
			if l < 0 {
				l = 0
			}
		}
		if tn.Right != nil {
			r = df(tn.Right)
			if r < 0 {
				r = 0
			}
		}

		ans = max(ans, l+r+val)
		val = max(l, r) + val

		return val
	}
	df(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
