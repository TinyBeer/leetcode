package main

/*
给定二叉树的根节点 root ，返回所有左叶子之和。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	ans := 0
	var dfs func(tn *TreeNode, isLeft bool)
	dfs = func(tn *TreeNode, isLeft bool) {
		if tn.Left != nil {
			dfs(tn.Left, true)
		}
		if tn.Right != nil {
			dfs(tn.Right, false)
		}
		if isLeft && tn.Left == nil && tn.Right == nil {
			ans += tn.Val
		}
	}
	dfs(root, false)
	return ans
}
