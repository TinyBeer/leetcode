package main

/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历 取队尾
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	var dp []*TreeNode
	dp = append(dp, root)

	for len(dp) > 0 {
		for i := len(dp) - 1; i >= 0; i-- {
			tr := dp[0]
			dp = dp[1:]
			if i == 0 {
				ans = append(ans, tr.Val)
			}
			if tr.Left != nil {
				dp = append(dp, tr.Left)
			}
			if tr.Right != nil {
				dp = append(dp, tr.Right)
			}
		}
	}
	return ans
}
