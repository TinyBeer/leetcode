package main

import (
	"strconv"
	"strings"
)

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	var tmp []string
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		tmp = append(tmp, strconv.Itoa(tn.Val))
		defer func() {
			tmp = tmp[:len(tmp)-1]
		}()
		if tn.Left == nil && tn.Right == nil {
			res = append(res, strings.Join(tmp, "->"))
			return
		}
		if tn.Left != nil {
			dfs(tn.Left)
		}
		if tn.Right != nil {
			dfs(tn.Right)
		}
	}
	dfs(root)
	return res
}
