package main

/*
给定一个二叉树 root ，返回其最大深度。

二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var df func(*TreeNode) int
	df = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}

		return max(df(tn.Left), df(tn.Right)) + 1
	}
	return df(root)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
