package main

/*
给定一个二叉树，判断它是否是
平衡二叉树

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var df func(*TreeNode) (int, bool)
	df = func(tn *TreeNode) (int, bool) {
		if tn == nil {
			return 0, true
		}
		lh, lb := df(tn.Left)
		rh, rb := df(tn.Right)
		if !lb || !rb || abs(lh-rh) >= 2 {
			return 0, false
		}
		return max(lh, rh) + 1, true
	}
	_, b := df(root)
	return b
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
