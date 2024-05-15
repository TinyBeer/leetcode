package main

import (
	"strconv"
	"strings"
)

/*
给定两个整数数组 preorder 和 inorder ，
其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	tbl := make(map[int]int, len(preorder))
	for i, v := range inorder {
		tbl[v] = i
	}

	var df func(p, l, r int) *TreeNode
	df = func(p, l, r int) *TreeNode {
		if p >= len(preorder) || l > r {
			return nil
		}
		i := tbl[preorder[p]]
		if preorder[p] == inorder[i] {
			return &TreeNode{
				Val:   preorder[p],
				Left:  df(p+1, l, i-1),
				Right: df(p+i-l+1, i+1, r),
			}
		}

		return nil
	}
	return df(0, 0, len(inorder)-1)

}
