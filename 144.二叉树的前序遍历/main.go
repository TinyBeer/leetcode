package main

/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var df func(*TreeNode)
	df = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		ans = append(ans, tn.Val)
		df(tn.Left)
		df(tn.Right)
	}
	df(root)
	return ans
}
