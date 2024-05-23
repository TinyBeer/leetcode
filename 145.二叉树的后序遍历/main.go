package main

/*
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var df func(*TreeNode)
	df = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.Left != nil {
			df(tn.Left)
		}
		if tn.Right != nil {
			df(tn.Right)
		}
		ans = append(ans, tn.Val)

	}
	df(root)
	return ans
}
