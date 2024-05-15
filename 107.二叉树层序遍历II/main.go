package main

import "slices"

/*
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。
 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var res [][]int
	for len(q) > 0 {
		layer := make([]int, 0)
		for size := len(q); size > 0; size-- {
			layer = append(layer, q[0].Val)
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			q = q[1:]
		}
		res = append(res, layer)
	}
	slices.Reverse(res)
	return res
}
