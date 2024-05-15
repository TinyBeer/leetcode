package main

import (
	"strconv"
	"strings"
)

/*
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵
平衡
 二叉搜索树。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOderTraversal(root *TreeNode) string {
	var df func(*TreeNode) []string
	df = func(tn *TreeNode) []string {
		if tn == nil {
			return []string{"null"}
		}
		return append(append([]string{strconv.Itoa(tn.Val)}, df(tn.Left)...), df(tn.Right)...)
	}
	return strings.Join(df(root), ",")
}

func sortedArrayToBST(nums []int) *TreeNode {
	var (
		dfs func(low, high int) *TreeNode
	)
	dfs = func(low, high int) *TreeNode {
		if low > high {
			return nil
		}
		mid := (low + high) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = dfs(low, mid-1)
		root.Right = dfs(mid+1, high)
		return root
	}
	return dfs(0, len(nums)-1)
}

// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	if len(nums) == 1 {
// 		return &TreeNode{Val: nums[0]}
// 	}
// 	mid := len(nums) >> 1

// 	return &TreeNode{
// 		Val:   nums[mid],
// 		Left:  sortedArrayToBST(nums[:mid]),
// 		Right: sortedArrayToBST(nums[mid+1:]),
// 	}
// }
