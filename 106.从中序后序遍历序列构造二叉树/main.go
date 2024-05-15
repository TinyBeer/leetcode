package main

import (
	"strconv"
	"strings"
)

/*
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历，
 postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	tbl := make(map[int]int, len(inorder))
	for i, v := range inorder {
		tbl[v] = i
	}

	var df func(p, l, r int) *TreeNode
	df = func(p, l, r int) *TreeNode {
		if p < 0 || l > r {
			return nil
		}
		i := tbl[postorder[p]]
		return &TreeNode{
			Val:   postorder[p],
			Left:  df(p+i-r-1, l, i-1),
			Right: df(p-1, i+1, r),
		}
	}

	return df(len(postorder)-1, 0, len(inorder)-1)
}
