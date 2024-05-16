package main

/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和
targetSum 。如果存在，返回 true ；否则，返回 false 。

叶子节点 是指没有子节点的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var df func(*TreeNode, int) bool
	df = func(tn *TreeNode, sum int) bool {
		if tn == nil {
			return false
		}
		sum += tn.Val
		if tn.Left == nil && tn.Right == nil && sum == targetSum {
			return true
		}
		return df(tn.Left, sum) || df(tn.Right, sum)
	}
	return df(root, 0)
}
