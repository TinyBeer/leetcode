package main

/*
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	var dp func(*TreeNode, int)
	dp = func(node *TreeNode, num int) {
		num = num*10 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += num
		}
		if node.Left != nil {
			dp(node.Left, num)
		}
		if node.Right != nil {
			dp(node.Right, num)
		}
	}
	dp(root, 0)
	return ans
}
